package h264

import (
	"fmt"
	"log"

	"github.com/kakami/pkg/bit"
)

var pixelAspect = [17]AVRational{
	{0, 1},
	{1, 1},
	{12, 11},
	{10, 11},
	{16, 11},
	{40, 33},
	{24, 11},
	{20, 11},
	{32, 11},
	{80, 33},
	{18, 11},
	{15, 11},
	{64, 33},
	{160, 99},
	{4, 3},
	{3, 2},
	{2, 1},
}

// AVCHeader ...
type AVCHeader struct {
	SPS

	AVCProfile  uint `json:"avc_profile"`
	AVCCompat   uint `json:"avc_compat"`
	AVCLevel    uint `json:"avc_level"`
	AVCNalBytes uint `json:"avc_nal_bytes"`
	NalSize     uint
	RbspSize    uint

	Width  int `json:"width"`
	Height int `json:"height"`
	Fps    int `json:"fps"`
}

// ParseAVCHeader ...
func (h *AVCHeader) ParseAVCHeader(in []byte) (err error) {
	fmt.Println("AVC:   ", in)
	r := bit.NewReader(in)
	r.Skip(48)
	h.AVCProfile = r.Read8()
	h.AVCCompat = r.Read8()
	h.AVCLevel = r.Read8()
	h.AVCNalBytes = r.Read8()&0x03 + 1

	// fmt.Printf("profile: %d, nal bytes: %d\n", h.AVCProfile, h.AVCNalBytes)

	// nnals
	nnals := r.Read8() & 0x1f
	// if r.Read8()&0x1f == 0 {
	if nnals == 0 {
		return fmt.Errorf("nnals is 0")
	}

	// nal size
	h.NalSize = uint(r.Read(16))

	// fmt.Printf("nnals: %d, nal size: %d, body size: %d\n", nnals, h.NalSize, len(in))

	// nal type
	if r.Read8() != 0x67 {
		return fmt.Errorf("nal is not 0x67")
	}

	// sps
	if err = h.ParseH264SPS(r); err != nil {
		log.Printf("parse sps failed, err: %s\n", err.Error())
		return
	}

	// fps
	// if h.NumUtilsInTick != 0 {
	// 	h.Fps = int(h.TimeScale / h.NumUtilsInTick / 2)
	// } else {
	// 	h.Fps = 25
	// }

	return
}

// ParseH264SPS ...
func (h *AVCHeader) ParseH264SPS(z *bit.Reader) (err error) {
	var rbsp []byte

	if rbsp, err = h.nalToRbsp(z); err != nil {
		return
	}

	r := bit.NewReader(rbsp)

	h.ProfileIdc = int(r.Read(8))
	h.ConstraintSetFlags |= int(r.Read(1)) << 0
	h.ConstraintSetFlags |= int(r.Read(1)) << 1
	h.ConstraintSetFlags |= int(r.Read(1)) << 2
	h.ConstraintSetFlags |= int(r.Read(1)) << 3
	h.ConstraintSetFlags |= int(r.Read(1)) << 4
	h.ConstraintSetFlags |= int(r.Read(1)) << 5

	r.Skip(2)
	h.LevelIdc = int(r.Read(8))
	h.SPSID = uint(r.ReadGolomb())
	if h.SPSID >= MAX_SPS_COUNT {
		return fmt.Errorf("h.264 sps: sps_id is too big")
	}

	h.TimeOffsetLength = 24
	h.FullRange = -1
	h.ScalingMatrixPresent = 0
	h.ColorSpace = AVCOL_SPC_UNSPECIFIED

	// fmt.Printf("sps$ level_idc: %d\n", h.LevelIdc)
	// fmt.Printf("sps$ sps_id: %d\n", h.SPSID)
	fmt.Printf("sps$ profile_idc: %d\n", h.ProfileIdc)

	if h.ProfileIdc == 100 || // High profile
		h.ProfileIdc == 110 || // High10 profile
		h.ProfileIdc == 122 || // High422 profile
		h.ProfileIdc == 244 || // High444 Predictive profile
		h.ProfileIdc == 44 || // Cavlc444 profile
		h.ProfileIdc == 83 || // Scalable Constrained High profile (SVC)
		h.ProfileIdc == 86 || // Scalable High Intra profile (SVC)
		h.ProfileIdc == 118 || // Stereo High profile (MVC)
		h.ProfileIdc == 128 || // Multiview High profile (MVC)
		h.ProfileIdc == 138 || // Multiview Depth High profile (MVCD)
		h.ProfileIdc == 144 { // old High444 profile
		h.ChromaFormatIdc = int(r.ReadGolomb())
		if h.ChromaFormatIdc > 3 {
			return fmt.Errorf("invalid chroma_format_idc: %d", h.ChromaFormatIdc)
		} else if h.ChromaFormatIdc == 3 {
			h.ResidualColorTransformFlag = int(r.Read(1))
			if h.ResidualColorTransformFlag != 0 {
				return fmt.Errorf("h.264 sps: separate color planes are not supported")
			}
		}

		h.BitDepthLuma = int(r.ReadGolomb()) + 8
		h.BitDepthChorma = int(r.ReadGolomb()) + 8
		if h.BitDepthChorma != h.BitDepthLuma {
			return fmt.Errorf("h.264 sps: bit_depth_luma != bit_depth_chorma")
		}

		if h.BitDepthLuma > 14 || h.BitDepthChorma > 14 {
			return fmt.Errorf("h.264 sps: illegal bit depth value(%d, %d)", h.BitDepthLuma, h.BitDepthChorma)
		}
		h.TransformBypass = int(r.Read(1))
		h.decodeScalingMatrices(r, 1)
	} else {
		h.ChromaFormatIdc = 1
		h.BitDepthLuma = 8
		h.BitDepthChorma = 8
	}

	log2MaxFrameNumMinus4 := r.ReadGolomb()
	if log2MaxFrameNumMinus4 < MIN_LOG2_MAX_FRAME_NUM-4 ||
		log2MaxFrameNumMinus4 > MAX_LOG2_MAX_FRAME_NUM-4 {
		return fmt.Errorf("h.264 sps: log2_max_frame_num_minus4 out of range(0-12): %d", log2MaxFrameNumMinus4)
	}
	h.Log2MaxFrameNum = int(log2MaxFrameNumMinus4) + 4
	h.PocType = int(r.ReadGolomb())

	if h.PocType == 0 {
		t := int(r.ReadGolomb())
		if t > 12 {
			return fmt.Errorf("h.264 sps: log2_max_poc_lsb (%d) is out of range", t)
		}
		h.Log2MaxPocLsb = t + 4
	} else if h.PocType == 1 {
		h.DeltaPicOrderAlwaysZeroFlag = int(r.Read(1))
		h.OffsetForNonRefPic = r.ReadSeGolomb()
		h.OffsetForTopToBottomField = r.ReadSeGolomb()
		h.PocCycleLength = int(r.ReadGolomb())

		if h.PocCycleLength >= len(h.OffsetForRefFrame) {
			return fmt.Errorf("poc_cycle_length overflow %d", h.PocCycleLength)
		}

		for i := 0; i < h.PocCycleLength; i++ {
			h.OffsetForRefFrame[i] = r.ReadSeGolomb()
		}
	} else if h.PocType != 2 {
		return fmt.Errorf("h.264 sps: illegal POC type %d", h.PocType)
	}

	if h.RefFrameCount = int(r.ReadGolomb()); h.RefFrameCount > 31 {
		return fmt.Errorf("h.264 sps: ref_frame_count is out of range: %d", h.RefFrameCount)
	}
	if h.RefFrameCount > H264_MAX_PICTURE_COUNT-2 || h.RefFrameCount > 16 {
		return fmt.Errorf("h.264 sps: too many reference frames %d", h.RefFrameCount)
	}

	h.GapsInFrameNumAllowedFlag = r.ReadInt(1)
	h.MbWidth = int(r.ReadGolomb()) + 1
	h.MbHeight = int(r.ReadGolomb()) + 1

	h.FrameMbsOnlyFlag = r.ReadInt(1)
	if h.FrameMbsOnlyFlag == 0 {
		h.MbAff = r.ReadInt(1)
	} else {
		h.MbAff = 0
	}
	h.Direct8x8InterfaceFlag = r.ReadInt(1)
	h.Crop = r.ReadInt(1)
	var cropLeft, cropRight, cropTop, cropBottom int
	if h.Crop != 0 {
		cropLeft = int(r.ReadGolomb())
		cropRight = int(r.ReadGolomb())
		cropTop = int(r.ReadGolomb())
		cropBottom = int(r.ReadGolomb())
	}
	h.Width = 16*h.MbWidth - (cropLeft+cropRight)*2
	h.Height = 16*h.MbHeight*(2-h.FrameMbsOnlyFlag) - (cropTop+cropBottom)*2

	h.VuiParametersPresentFlag = r.ReadInt(1)
	if h.VuiParametersPresentFlag != 0 {
		if err = h.decodeVuiParameters(r); err != nil {
			return
		}
	}

	if h.Sar.Den == 0 {
		h.Sar.Den = 1
	}

	return
}

func (h *AVCHeader) decodeVuiParameters(r *bit.Reader) (err error) {
	aspectRatioInfoPresentFlag := r.ReadInt(1)
	var aspectRatioIdc int
	if aspectRatioInfoPresentFlag != 0 {
		aspectRatioIdc = r.ReadInt(8)
		if aspectRatioIdc == EXTENDED_SAR {
			h.Sar.Num = r.ReadInt(16)
			h.Sar.Den = r.ReadInt(16)
		} else if aspectRatioIdc < len(pixelAspect) {
			h.Sar = pixelAspect[aspectRatioIdc]
		} else {
			return fmt.Errorf("h.264 sps: aspect_ratio_idc is out of range: %d", aspectRatioIdc)
		}
	} else {
		h.Sar.Num = 0
		h.Sar.Den = 0
	}

	if r.Read(1) != 0 { // overscan_info_present_flag
		r.Read(1) // overscan_appropriate_flag
	}

	if h.VideoSignalTypePresentFlag = r.ReadInt(1); h.VideoSignalTypePresentFlag != 0 {
		r.Read(3)                  // video_format
		h.FullRange = r.ReadInt(1) // video_full_range_flag
		if h.ColourDescriptionPresentFlag = r.ReadInt(1); h.ColourDescriptionPresentFlag != 0 {
			h.ColorPrimaries = r.ReadUInt(8)
			h.ColorTrc = r.ReadUInt(8)
			h.ColorSpace = r.ReadUInt(8)
			if h.ColorPrimaries >= AVCOL_PRI_NB {
				h.ColorPrimaries = AVCOL_PRI_UNSPECIFIED
			}
			if h.ColorTrc >= AVCOL_TRC_NB {
				h.ColorTrc = AVCOL_TRC_UNSPECIFIED
			}
			if h.ColorSpace >= AVCOL_SPC_NB {
				h.ColorSpace = AVCOL_SPC_UNSPECIFIED
			}
		}
	}

	if r.Read(1) != 0 {
		r.ReadGolomb() // chroma_sample_location_type_top_field
		r.ReadGolomb() // chroma_sample_location_type_bottom_field
	}

	if r.Read(1) != 0 && r.Left() < 10 {
		return
	}

	if h.TimingInfoPresentFlag = r.ReadInt(1); h.TimingInfoPresentFlag != 0 {
		h.NumUtilsInTick = uint32(r.Read(32))
		h.TimeScale = uint32(r.Read(32))
		if h.NumUtilsInTick == 0 {
			return fmt.Errorf("num_units_in_tick is 0")
		}
		if h.TimeScale == 0 {
			return fmt.Errorf("time_scale is 0")
		}
		h.FixedFrameRateFlag = r.ReadInt(1)
	}

	if h.NalHrdParametersPresentFlag = r.ReadInt(1); h.NalHrdParametersPresentFlag != 0 {
		if err = h.decodeHdrParameters(r); err != nil {
			return
		}
	}
	if h.VclHrdParametersPresentFlag = r.ReadInt(1); h.VclHrdParametersPresentFlag != 0 {
		if err = h.decodeHdrParameters(r); err != nil {
			return
		}
	}
	if h.VclHrdParametersPresentFlag != 0 || h.NalHrdParametersPresentFlag != 0 {
		r.Read(1)
	}
	h.PicStructPresentFlag = r.ReadInt(1)
	if r.Left() <= 0 {
		return
	}
	if h.BitStreamRestrictionFlag = r.ReadInt(1); h.BitStreamRestrictionFlag != 0 {
		r.Read(1)      // motion_vectors_over_pic_boundaries_flag
		r.ReadGolomb() // max_bytes_per_pic_denom
		r.ReadGolomb() // max_bits_per_mb_denom
		r.ReadGolomb() // log2_max_mv_length_horizontal
		r.ReadGolomb() // log2_max_mv_length_vertical
		h.NumReorderFrames = int(r.ReadGolomb())
		r.ReadGolomb() // max_dec_frame_buffering

		if r.Left() <= 0 {
			h.NumReorderFrames = 0
			h.BitStreamRestrictionFlag = 0
		}

		if h.NumReorderFrames > 16 {
			err = fmt.Errorf("h.264 sps: num_reorder_frames is out of range: %d", h.NumReorderFrames)
			h.NumReorderFrames = 16
			return
		}
	}

	if r.Left() <= 0 {
		return fmt.Errorf("h.264 sps: more data needed")
	}

	return
}

func (h *AVCHeader) decodeHdrParameters(r *bit.Reader) (err error) {
	cpbCount := int(r.ReadGolomb()) + 1
	if cpbCount > 32 {
		return fmt.Errorf("cpb_count is out of range: %d", cpbCount)
	}
	r.Read(4) // bit_rate_scale
	r.Read(4) // cpb_size_scale
	for i := 0; i < cpbCount; i++ {
		r.ReadGolomb() // bit_rate_value_minus1
		r.ReadGolomb() // cpb_size_value_minus1
		r.Read(1)      // cbr_flag
	}
	h.InitialCpbRemovalDelayLength = r.ReadInt(5) + 1
	h.CpbRemovalDelayLength = r.ReadInt(5) + 1
	h.DpbOutputDelayLength = r.ReadInt(5) + 1
	h.TimeOffsetLength = r.ReadInt(5) + 1
	h.CpbCnt = cpbCount
	return
}

func (h *AVCHeader) decodeScalingList16(r *bit.Reader, sm *[16]uint8, jvtList *[16]uint8, fallbackList *[16]uint8) {
	var size = 16

	if r.Read(1) == 0 {
		for i := 0; i < size; i++ {
			sm[i] = zigzagScan[i]
		}
	} else {
		last := 8
		next := 8
		for i := 0; i < size; i++ {
			if next != 0 {
				next = last + int(r.ReadSeGolomb())&0xff
			}
			if i == 0 && next == 0 {
				for j := 0; j < size; j++ {
					sm[j] = zigzagScan[j]
				}
				break
			}
			if sm[zigzagScan[i]] == uint8(next) {
				last = next
			}
		}
	}
}

func (h *AVCHeader) decodeScalingList64(r *bit.Reader, sm *[64]uint8, jvtList *[64]uint8, fallbackList *[64]uint8) {
	var size = 64

	if r.Read(1) == 0 {
		for i := 0; i < size; i++ {
			sm[i] = ffZigzagDirect[i]
		}
	} else {
		last := 8
		next := 8
		for i := 0; i < size; i++ {
			if next != 0 {
				next = last + int(r.ReadSeGolomb())&0xff
			}
			if i == 0 && next == 0 {
				for j := 0; j < size; j++ {
					sm[j] = ffZigzagDirect[j]
				}
				break
			}
			if sm[ffZigzagDirect[i]] == uint8(next) {
				last = next
			}
		}
	}
}

func (h *AVCHeader) decodeScalingMatrices(r *bit.Reader, isSPS int) {
	fallbackSPS := isSPS == 0 && h.ScalingMatrixPresent != 0
	var fallback16 [2][16]uint8
	var fallback64 [2][64]uint8

	if fallbackSPS {
		fallback16[0] = h.ScalingMatrix4[0]
		fallback16[1] = h.ScalingMatrix4[3]
		fallback64[0] = h.ScalingMatrix8[0]
		fallback64[1] = h.ScalingMatrix8[3]
	} else {
		fallback16[0] = defaultScaling4[0]
		fallback16[1] = defaultScaling4[1]
		fallback64[0] = defaultScaling8[0]
		fallback64[1] = defaultScaling8[1]
	}

	if r.Read(1) != 0 {
		h.ScalingMatrixPresent |= isSPS
		h.decodeScalingList16(r, &h.ScalingMatrix4[0], &defaultScaling4[0], &fallback16[0])
		h.decodeScalingList16(r, &h.ScalingMatrix4[1], &defaultScaling4[0], &h.ScalingMatrix4[0])
		h.decodeScalingList16(r, &h.ScalingMatrix4[2], &defaultScaling4[0], &h.ScalingMatrix4[1])
		h.decodeScalingList16(r, &h.ScalingMatrix4[3], &defaultScaling4[1], &fallback16[1])
		h.decodeScalingList16(r, &h.ScalingMatrix4[4], &defaultScaling4[1], &h.ScalingMatrix4[3])
		h.decodeScalingList16(r, &h.ScalingMatrix4[5], &defaultScaling4[1], &h.ScalingMatrix4[4])
		if isSPS != 0 {
			h.decodeScalingList64(r, &h.ScalingMatrix8[0], &defaultScaling8[0], &fallback64[0])
			h.decodeScalingList64(r, &h.ScalingMatrix8[3], &defaultScaling8[1], &fallback64[1])
			if h.ChromaFormatIdc == 3 {
				h.decodeScalingList64(r, &h.ScalingMatrix8[1], &defaultScaling8[0], &h.ScalingMatrix8[0])
				h.decodeScalingList64(r, &h.ScalingMatrix8[4], &defaultScaling8[1], &h.ScalingMatrix8[3])
				h.decodeScalingList64(r, &h.ScalingMatrix8[2], &defaultScaling8[0], &h.ScalingMatrix8[1])
				h.decodeScalingList64(r, &h.ScalingMatrix8[5], &defaultScaling8[1], &h.ScalingMatrix8[4])
			}
		}
	}
}

func (h *AVCHeader) nalToRbsp(r *bit.Reader) (rbsp []byte, err error) {
	var i, j, count, rbspSize uint
	var nal []byte

	nalSize := h.NalSize
	rbspSize = nalSize + 4

	if nal, err = r.GetSliceCopy(nalSize); err != nil {
		return
	}

	for i = 0; i < nalSize; i++ {
		if count == 2 && nal[i] < 0x03 {
			err = fmt.Errorf("n NAL unit, 0x000000, 0x000001 or 0x000002 shall not occur at any byte-aligned position")
			return
		}
		if count == 2 && nal[i] == 0x03 {
			if i < nalSize-1 && nal[i+1] > 0x03 {
				return
			}

			if i == nalSize-1 {
				break
			}
			i++
			count = 0
		}

		if j >= rbspSize {
			err = fmt.Errorf("not enough space")
			return
		}
		rbsp = append(rbsp, nal[i])
		if nal[i] == 0x00 {
			count++
		} else {
			count = 0
		}
		j++
	}

	h.NalSize = i
	h.RbspSize = j

	return
}
