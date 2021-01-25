package h265

import (
	"fmt"

	"github.com/kakami/pkg/bit"
)

// NalUnit ...
type NalUnit struct {
	NalUnitLength uint   `json:"nalUnitLength"`
	NalUnitData   []byte `json:"nalUnit"`
}

// Nal ...
type Nal struct {
	ArrayCompleteness uint      `json:"array_completeness"`
	NalUnitType       uint      `json:"NAL_unit_type"`
	NumNalus          uint      `json:"numNalus"`
	NalUnits          []NalUnit `json:"nal_units"`
}

// HEVCHeader ...
type HEVCHeader struct {
	Width                            uint   `json:"width"`
	Height                           uint   `json:"height"`
	ConfigurationVersion             uint   `json:"configurationVersion"`
	GeneralProfileSpace              uint   `json:"general_profile_space"`
	GeneralTierFlag                  uint   `json:"general_tier_flag"`
	GeneralProfileIdc                uint   `json:"general_profile_idc"`
	GeneralProfileCompatibilityFlags uint   `json:"general_profile_compatibility_flags"`
	GeneralConstraintIndicatorFlags  uint64 `json:"general_constraint_indicator_flags"`
	GeneralLevelIdc                  uint   `json:"general_level_idc"`
	MinSpatialSegmentationIdc        uint   `json:"min_spatial_segmentation_idc"`
	ParallelismType                  uint   `json:"parallelism_type"`
	ChromaFormatIdc                  uint   `json:"chroma_format_idc"`
	BitDepthLumaMinus8               uint   `json:"bit_depth_luma_minus8"`
	BitDepthChromaMinus8             uint   `json:"bit_depth_chroma_minus8"`
	AvgFrameRate                     uint   `json:"avgFrameRate"`
	ConstantFrameRate                uint   `json:"constantFrameRate"`
	NumTemporalLayers                uint   `json:"numTemporalLayers"`
	TemporalIDNested                 uint   `json:"temporalIdNested"`
	LengthSizeMinusOne               uint   `json:"lengthSizeMinusOne"`
	NumOfArrays                      uint   `json:"numOfArrays"`
	Nals                             []Nal
	SPSs                             []SPS
}

// ParseHEVCHeader ...
func (h *HEVCHeader) ParseHEVCHeader(in []byte) (err error) {
	r := bit.NewReader(in)
	r.Skip(40)
	h.ConfigurationVersion = r.Read8()
	h.GeneralProfileSpace = r.ReadUInt(2)
	h.GeneralTierFlag = r.ReadUInt(1)
	h.GeneralProfileIdc = r.ReadUInt(5)
	h.GeneralProfileCompatibilityFlags = r.ReadUInt(32)
	h.GeneralConstraintIndicatorFlags = r.Read(48)
	h.GeneralLevelIdc = r.ReadUInt(8)
	r.Skip(4)
	h.MinSpatialSegmentationIdc = r.ReadUInt(12)
	r.Skip(6)
	h.ParallelismType = r.ReadUInt(2)
	r.Skip(6)
	h.ChromaFormatIdc = r.ReadUInt(2)
	r.Skip(5)
	h.BitDepthLumaMinus8 = r.ReadUInt(3)
	r.Skip(5)
	h.BitDepthChromaMinus8 = r.ReadUInt(3)

	h.AvgFrameRate = r.ReadUInt(16)
	h.ConstantFrameRate = r.ReadUInt(2)
	h.NumTemporalLayers = r.ReadUInt(3)
	h.TemporalIDNested = r.ReadUInt(1)
	h.LengthSizeMinusOne = r.ReadUInt(2)
	h.NumOfArrays = r.ReadUInt(8)

	var i, j uint
	for i = 0; i < h.NumOfArrays; i++ {
		h.Nals = append(h.Nals, Nal{})
		h.Nals[i].ArrayCompleteness = r.ReadUInt(1)
		r.Skip(1)
		h.Nals[i].NalUnitType = r.ReadUInt(6)
		h.Nals[i].NumNalus = r.ReadUInt(16)
		for j = 0; j < h.Nals[i].NumNalus; j++ {
			h.Nals[i].NalUnits = append(h.Nals[i].NalUnits, NalUnit{})
			h.Nals[i].NalUnits[j].NalUnitLength = r.ReadUInt(16)
			s, err := r.GetSliceCopy(h.Nals[i].NalUnits[j].NalUnitLength)
			if err != nil {
				return err
			}
			if h.Nals[i].NalUnitType == NAL_SPS {
				h.Nals[i].NalUnits[j].NalUnitData = append(h.Nals[i].NalUnits[j].NalUnitData, s...)
			}
		}
	}

	for i = 0; i < h.NumOfArrays; i++ {
		if h.Nals[i].NalUnitType == NAL_SPS {
			for j = 0; j < h.Nals[i].NumNalus; j++ {
				if err = h.parseHEVCSPS(h.Nals[i].NalUnits[j].NalUnitData); err != nil {
					return
				}
			}
		}
	}

	return
}

func (h *HEVCHeader) parseHEVCSPS(nal []byte) (err error) {
	var rbsp []byte
	var i uint

	if rbsp, err = h.nalToRbsp(nal); err != nil {
		return
	}
	h.SPSs = append(h.SPSs, SPS{})
	sps := &h.SPSs[len(h.SPSs)-1]
	sps.NalUnitType = NAL_SPS
	r := bit.NewReader(rbsp)
	r.Skip(16)

	sps.SPSSeqParameterSetID = r.ReadUInt(4)
	sps.SPSMaxSubLayersMinus1 = r.ReadUInt(3)
	sps.SPSTemporalIDNestingFlag = r.ReadUInt(1)

	h.parseProfileTierLevel(r, sps)

	sps.SPSSeqParameterSetID = uint(r.ReadGolomb())
	if sps.SPSSeqParameterSetID > 16 || r.Left() <= 0 {
		return fmt.Errorf("read sps_seq_parameter_set_id error: %d", sps.SPSSeqParameterSetID)
	}

	sps.ChromaFormatIdc = uint(r.ReadGolomb())
	if sps.ChromaFormatIdc > 3 || r.Left() <= 0 {
		return fmt.Errorf("read chroma_format_idc error: %d", sps.ChromaFormatIdc)
	}

	if sps.ChromaFormatIdc == 3 {
		sps.SeparateColourPlaneFlag = r.ReadUInt(1)
	}

	sps.PicWidthInLumaSamples = uint(r.ReadGolomb())
	sps.PicHeightInLumaSamples = uint(r.ReadGolomb())
	h.Width = sps.PicWidthInLumaSamples
	h.Height = sps.PicHeightInLumaSamples

	if sps.ConformanceWindowFlag = r.ReadUInt(1); sps.ConformanceWindowFlag > 0 {
		sps.ConfWinLeftOffset = uint(r.ReadGolomb())
		sps.ConfWinRightOffset = uint(r.ReadGolomb())
		sps.ConfWinTopOffset = uint(r.ReadGolomb())
		sps.ConfWinBottomOffset = uint(r.ReadGolomb())

		var subWidthC, subHeightC uint
		if sps.ChromaFormatIdc == 1 {
			subWidthC = 2
			subHeightC = 2
		} else if sps.ChromaFormatIdc == 2 {
			subWidthC = 2
			subHeightC = 1
		} else {
			subWidthC = 1
			subHeightC = 1
		}
		h.Width = sps.PicWidthInLumaSamples - (subWidthC*sps.ConfWinRightOffset + 1) - (subWidthC * sps.ConfWinLeftOffset)
		h.Height = sps.PicHeightInLumaSamples - (subHeightC*sps.ConfWinBottomOffset + 1) - (subHeightC * sps.ConfWinTopOffset)
	}

	sps.BitDepthLumaMinus8 = uint(r.ReadGolomb())
	sps.BitDepthChromaMinus8 = uint(r.ReadGolomb())
	sps.Log2MaxPicOrderCntLsbMinus4 = uint(r.ReadGolomb())
	sps.SPSSubLayerOrderingInfoPresentFlag = r.ReadUInt(1)

	i = sps.SPSMaxSubLayersMinus1
	if sps.SPSSubLayerOrderingInfoPresentFlag > 0 {
		i = 0
	}
	for j := 0; i <= sps.SPSMaxSubLayersMinus1; i++ {
		sps.SPSMaxDecPicBufferingMinus1[j] = uint8(r.ReadGolomb())
		sps.SPSMaxNumReorderPics[j] = uint8(r.ReadGolomb())
		sps.SPSMaxLatencyIncreasePlus1[j] = uint8(r.ReadGolomb())
		j++
	}
	sps.Log2MinLumaCodingBlockSizeMinus3 = uint(r.ReadGolomb())
	sps.Log2DiffMaxMinLumaCodingBlockSize = uint(r.ReadGolomb())
	sps.Log2MinLumaTransformBlockSizeMinus2 = uint(r.ReadGolomb())
	sps.Log2DiffMaxMinLumaTransformBlockSize = uint(r.ReadGolomb())
	sps.MaxTransformHierarchyDepthInter = uint(r.ReadGolomb())
	sps.MaxTransformHierarchyDepthIntra = uint(r.ReadGolomb())

	sps.ScalingListEnabledFlag = r.ReadUInt(1)
	sps.SPSScalingListDataPresentFlag = r.ReadUInt(1)
	if sps.ScalingListEnabledFlag > 0 && sps.SPSScalingListDataPresentFlag > 0 {
		h.skipScalingData(r, sps)
	}

	sps.AmpEnabledFlag = r.ReadUInt(1)
	sps.SampleAdaptiveOffsetEnabledFlag = r.ReadUInt(1)
	if sps.PcmEnabledFlag = r.ReadUInt(1); sps.PcmEnabledFlag > 0 {
		sps.PcmSampleBitDepthLumaMinus1 = r.ReadUInt(4)
		sps.PcmSampleBitDepthChromaMinus1 = r.ReadUInt(4)
		sps.Log2MinPcmLumaCodingBlockSizeMinus3 = uint(r.ReadGolomb())
		sps.Log2DiffMaxMinLumaCodingBlockSize = uint(r.ReadGolomb())
		sps.PcmLoopFilterDisabledFlag = r.ReadUInt(1)
	}

	if sps.NumShortTermRefPicSets = uint(r.ReadGolomb()); sps.NumShortTermRefPicSets > HEVC_MAX_SHORT_TERM_REF_PIC_SETS {
		return fmt.Errorf("read num_short_term_ref_pic_sets error: %d", sps.NumShortTermRefPicSets)
	}
	for i = 0; i < sps.NumShortTermRefPicSets; i++ {
		if err = h.parseRps(r, i, sps); err != nil {
			return
		}
	}

	if sps.LongTermRefPicsPresentFlag = r.ReadUInt(1); sps.LongTermRefPicsPresentFlag != 0 {
		if sps.NumLongTermRefPicsSPS = uint(r.ReadGolomb()); sps.NumLongTermRefPicsSPS > 31 {
			return fmt.Errorf("read num_long_term_ref_pics_sps error: %d", sps.NumLongTermRefPicsSPS)
		}
		for i = 0; i < sps.NumLongTermRefPicsSPS; i++ {
			len := sps.Log2MaxPicOrderCntLsbMinus4 + 4
			if len > 16 {
				len = 16
			}
			sps.LtRefPicPocLsbSPS[i] = uint16(r.Read(uint32(len)))
			sps.UsedByCurrPicLtSPSFlag[i] = uint8(r.Read(1))
		}
	}

	sps.SPSTemporalMvpEnabledFlag = r.ReadUInt(1)
	sps.StrongIntraSmoothingEnabledFlag = r.ReadUInt(1)

	if sps.VuiParametersPresentFlag = r.ReadUInt(1); sps.VuiParametersPresentFlag != 0 {
		h.parseHvccVui()
	}

	return
}

func (h *HEVCHeader) parseProfileTierLevel(r *bit.Reader, sps *SPS) (err error) {
	ptl := &sps.PTL
	ptl.GeneralProfileSpace = r.ReadUInt(2)
	ptl.GeneralTierFlag = r.ReadUInt(1)
	ptl.GeneralProfileIdc = r.ReadUInt(5)
	for i := 0; i < 32; i++ {
		ptl.GeneralProfileCompatibilityFlag[i] = uint8(r.ReadUInt(1))
	}
	ptl.GeneralProgressiveSourceFlag = r.ReadUInt(1)
	ptl.GeneralInterlacedSourceFlag = r.ReadUInt(1)
	ptl.GeneralNonPackedConstraintFlag = r.ReadUInt(1)
	ptl.GeneralFrameOnlyConstraintFlag = r.ReadUInt(1)

	ptl.GeneralMax12bitConstraintFlag = r.ReadUInt(1)
	ptl.GeneralMax10bitConstraintFlag = r.ReadUInt(1)
	ptl.GeneralMax8bitConstraintFlag = r.ReadUInt(1)
	ptl.GeneralMax422chromaConstraintFlag = r.ReadUInt(1)
	ptl.GeneralMax420chromaConstraintFlag = r.ReadUInt(1)
	ptl.GeneralMaxMonochromeConstraintFlag = r.ReadUInt(1)
	ptl.GeneralIntraConstraintFlag = r.ReadUInt(1)
	ptl.GeneralOnePictureOnlyConstraintFlag = r.ReadUInt(1)
	ptl.GeneralLowerBitRateConstraintFlag = r.ReadUInt(1)
	ptl.GeneralMax14bitConstraintFlag = r.ReadUInt(1)
	r.Skip(33)
	ptl.GeneralInbldFlag = r.ReadUInt(1)
	ptl.GeneralLevelIdc = r.ReadUInt(8)

	var i uint
	for i = 0; i < sps.SPSMaxSubLayersMinus1; i++ {
		ptl.SubLayerProfilePresentFlag[i] = uint8(r.ReadUInt(1))
		ptl.SubLayerLevelPresentFlag[i] = uint8(r.ReadUInt(1))
	}
	if sps.SPSMaxSubLayersMinus1 > 0 {
		for ; i < 8; i++ {
			r.Skip(2)
		}
	}

	for i = 0; i < sps.SPSMaxSubLayersMinus1; i++ {
		if ptl.SubLayerProfilePresentFlag[i] != 0 {
			ptl.SubLayerProfileSpace[i] = uint8(r.ReadUInt(2))
			ptl.SubLayerTierFlag[i] = uint8(r.ReadUInt(1))
			ptl.SubLayerProfileIdc[i] = uint8(r.ReadUInt(5))
			for j := 0; j < 32; j++ {
				ptl.SubLayerProfileCompatibilityFlag[i][j] = uint8(r.ReadUInt(1))
			}
			ptl.SubLayerProgressiveSourceFlag[i] = uint8(r.ReadUInt(1))
			ptl.SubLayerInterlacedSourceFlag[i] = uint8(r.ReadUInt(1))
			ptl.SubLayerNonPackedConstraintFlag[i] = uint8(r.ReadUInt(1))
			ptl.SubLayerFrameOnlyConstraintFlag[i] = uint8(r.ReadUInt(1))
			r.Skip(44)
		}
		if ptl.SubLayerLevelPresentFlag[i] != 0 {
			r.Skip(8)
		}
	}
	return
}

func (h *HEVCHeader) skipScalingData(r *bit.Reader, sps *SPS) (err error) {
	var i, j, k, t, numCoeffs uint
	sl := &sps.ScalingList
	for i = 0; i < 4; i++ {
		t = 6
		if i == 3 {
			t = 2
		}
		for j = 0; j < t; j++ {
			if sl.PredModeFlag[i][j] = uint8(r.Read(1)); sl.PredModeFlag[i][j] > 0 {
				sl.PredMatrixIDDelta[i][j] = uint8(r.ReadGolomb())
			} else {
				numCoeffs = 1 << (4 + (i << 1))
				if numCoeffs > 64 {
					numCoeffs = 64
				}
				if i > 1 {
					sl.DcCoefMinus8[i-2][j] = int16(r.ReadSeGolomb())
				}
				for k = 0; k < numCoeffs; k++ {
					sl.DeltaCoeff[i][j][k] = int8(r.ReadSeGolomb())
				}
			}
		}
	}
	return
}

func (h *HEVCHeader) parseRps(r *bit.Reader, idx uint, sps *SPS) (err error) {
	var i uint
	st := &sps.STRefPicSet[idx]
	st.InterRefPicSetPredictionFlag = uint8(r.Read(1))
	if idx > 0 && st.InterRefPicSetPredictionFlag > 0 {
		if idx >= sps.NumShortTermRefPicSets {
			return fmt.Errorf("invalid data, rps index overflow")
		}
		st.DeltaRpsSign = uint8(r.Read(1))
		st.AbsDeltaRpsMinus1 = uint16(r.ReadGolomb())
		sps.NumDeltaPocs[idx] = 0

		for i = 0; i <= sps.NumDeltaPocs[idx-1]; i++ {
			if st.UsedByCurrPicFlag[i] = uint8(r.Read(1)); st.UsedByCurrPicFlag[i] == 0 {
				st.UseDeltaFlag[i] = uint8(r.Read(1))
			}
			if st.UsedByCurrPicFlag[i] != 0 || st.UseDeltaFlag[i] != 0 {
				sps.NumDeltaPocs[idx]++
			}
		}
	} else {
		st.NumNegativePics = uint8(r.ReadGolomb())
		st.NumPositivePics = uint8(r.ReadGolomb())
		sps.NumDeltaPocs[idx] = uint(st.NumNegativePics + st.NumPositivePics)
		if (st.NumPositivePics+st.NumNegativePics)*2 > uint8(r.Left()) {
			return fmt.Errorf("invalid data, need %d, %d left", (st.NumPositivePics+st.NumNegativePics)*2, r.Left())
		}
		for i = 0; i < uint(st.NumNegativePics); i++ {
			st.DeltaPocS0Minus1[i] = uint16(r.ReadGolomb())
			st.UsedByCurrPicS0Flag[i] = uint8(r.Read(1))
		}
		for i = 0; i < uint(st.NumPositivePics); i++ {
			st.DeltaPocS1Minus1[i] = uint16(r.ReadGolomb())
			st.UsedByCurrPicS1Flag[i] = uint8(r.ReadGolomb())
		}
	}
	return
}

func (h *HEVCHeader) parseHvccVui() (err error) {
	return
}

func (h *HEVCHeader) nalToRbsp(nal []byte) (rbsp []byte, err error) {
	r := bit.NewReader(nal)
	r.Skip(1)
	if r.ReadInt(6) != NAL_SPS {
		err = fmt.Errorf("nal unit type is not sps")
		return
	}
	r.Skip(9)

	var count uint
	for i := 0; i < len(nal); i++ {
		if count == 2 {
			if nal[i] < 0x03 {
				err = fmt.Errorf("three bytes sequence error")
				return
			}
			if nal[i] == 0x03 && nal[i+1] > 0x03 {
				err = fmt.Errorf("four bytes sequence error")
				return
			}

			if nal[i] == 0x03 {
				count = 0
				continue
			}
		}

		rbsp = append(rbsp, nal[i])
		if nal[i] == 0x00 {
			count++
		} else {
			count = 0
		}
	}
	return
}
