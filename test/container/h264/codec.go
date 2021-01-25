// some codec valud

package h264

const (
	// hevc nal type
	NAL_TRAIL_N    = 0
	NAL_TRAIL_R    = 1
	NAL_TSA_N      = 2
	NAL_TSA_R      = 3
	NAL_STSA_N     = 4
	NAL_STSA_R     = 5
	NAL_RADL_N     = 6
	NAL_RADL_R     = 7
	NAL_RASL_N     = 8
	NAL_RASL_R     = 9
	NAL_BLA_W_LP   = 16
	NAL_BLA_W_RADL = 17
	NAL_BLA_N_LP   = 18
	NAL_IDR_W_RADL = 19
	NAL_IDR_N_LP   = 20
	NAL_CRA_NUT    = 21
	NAL_VPS        = 32
	NAL_SPS        = 33
	NAL_PPS        = 34
	NAL_AUD        = 35
	NAL_EOS_NUT    = 36
	NAL_EOB_NUT    = 37
	NAL_FD_NUT     = 38
	NAL_SEI_PREFIX = 39
	NAL_SEI_SUFFIX = 40

	// AVColorPrimaries
	AVCOL_PRI_BT709       = 1 ///< also ITU-R BT1361 / IEC 61966-2-4 / SMPTE RP177 Annex B
	AVCOL_PRI_UNSPECIFIED = 2
	AVCOL_PRI_RESERVED    = 3
	AVCOL_PRI_BT470M      = 4
	AVCOL_PRI_BT470BG     = 5 ///< also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM
	AVCOL_PRI_SMPTE170M   = 6 ///< also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC
	AVCOL_PRI_SMPTE240M   = 7 ///< functionally identical to above
	AVCOL_PRI_FILM        = 8
	AVCOL_PRI_BT2020      = 9  ///< ITU-R BT2020
	AVCOL_PRI_NB          = 10 ///< Not part of ABI

	// Color Transfer Characteristic.
	AVCOL_TRC_BT709        = 1 ///< also ITU-R BT1361
	AVCOL_TRC_UNSPECIFIED  = 2
	AVCOL_TRC_RESERVED     = 3
	AVCOL_TRC_GAMMA22      = 4 ///< also ITU-R BT470M / ITU-R BT1700 625 PAL & SECAM
	AVCOL_TRC_GAMMA28      = 5 ///< also ITU-R BT470BG
	AVCOL_TRC_SMPTE170M    = 6 ///< also ITU-R BT601-6 525 or 625 / ITU-R BT1358 525 or 625 / ITU-R BT1700 NTSC
	AVCOL_TRC_SMPTE240M    = 7
	AVCOL_TRC_LINEAR       = 8  ///< "Linear transfer characteristics"
	AVCOL_TRC_LOG          = 9  ///< "Logarithmic transfer characteristic (100:1 range)"
	AVCOL_TRC_LOG_SQRT     = 10 ///< "Logarithmic transfer characteristic (100 * Sqrt(10) : 1 range)"
	AVCOL_TRC_IEC61966_2_4 = 11 ///< IEC 61966-2-4
	AVCOL_TRC_BT1361_ECG   = 12 ///< ITU-R BT1361 Extended Colour Gamut
	AVCOL_TRC_IEC61966_2_1 = 13 ///< IEC 61966-2-1 (sRGB or sYCC)
	AVCOL_TRC_BT2020_10    = 14 ///< ITU-R BT2020 for 10 bit system
	AVCOL_TRC_BT2020_12    = 15 ///< ITU-R BT2020 for 12 bit system
	AVCOL_TRC_NB           = 16 ///< Not part of ABI

	// YUV colorspace type.
	AVCOL_SPC_RGB         = 0
	AVCOL_SPC_BT709       = 1 ///< also ITU-R BT1361 / IEC 61966-2-4 xvYCC709 / SMPTE RP177 Annex B
	AVCOL_SPC_UNSPECIFIED = 2
	AVCOL_SPC_RESERVED    = 3
	AVCOL_SPC_FCC         = 4
	AVCOL_SPC_BT470BG     = 5 ///< also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM / IEC 61966-2-4 xvYCC601
	AVCOL_SPC_SMPTE170M   = 6 ///< also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC / functionally identical to above
	AVCOL_SPC_SMPTE240M   = 7
	AVCOL_SPC_YCOCG       = 8  ///< Used by Dirac / VC-2 and H.264 FRext see ITU-T SG16
	AVCOL_SPC_BT2020_NCL  = 9  ///< ITU-R BT2020 non-constant luminance system
	AVCOL_SPC_BT2020_CL   = 10 ///< ITU-R BT2020 constant luminance system
	AVCOL_SPC_NB          = 11 ///< Not part of ABI

	MAX_SPS_COUNT          = 32
	MAX_LOG2_MAX_FRAME_NUM = 16
	MIN_LOG2_MAX_FRAME_NUM = 4
	// MKTAG = (a,b,c,d) ((a) | ((b) << 8) | ((c) << 16) | ((unsigned)(d) << 24))
	H264_MAX_PICTURE_COUNT = 36

	EXTENDED_SAR = 255
)

var (
	ffZigzagDirect = [64]uint8{
		0, 1, 8, 16, 9, 2, 3, 10,
		17, 24, 32, 25, 18, 11, 4, 5,
		12, 19, 26, 33, 40, 48, 41, 34,
		27, 20, 13, 6, 7, 14, 21, 28,
		35, 42, 49, 56, 57, 50, 43, 36,
		29, 22, 15, 23, 30, 37, 44, 51,
		58, 59, 52, 45, 38, 31, 39, 46,
		53, 60, 61, 54, 47, 55, 62, 63,
	}

	zigzagScan = [16]uint8{
		0 + 0*4, 1 + 0*4, 0 + 1*4, 0 + 2*4,
		1 + 1*4, 2 + 0*4, 3 + 0*4, 2 + 1*4,
		1 + 2*4, 0 + 3*4, 1 + 3*4, 2 + 2*4,
		3 + 1*4, 3 + 2*4, 2 + 3*4, 3 + 3*4,
	}

	defaultScaling4 = [2][16]uint8{
		{6, 13, 20, 28, 13, 20, 28, 32, 20, 28, 32, 37, 28, 32, 37, 42},
		{10, 14, 20, 24, 14, 20, 24, 27, 20, 24, 27, 30, 24, 27, 30, 34},
	}

	defaultScaling8 = [2][64]uint8{
		{
			6, 10, 13, 16, 18, 23, 25, 27,
			10, 11, 16, 18, 23, 25, 27, 29,
			13, 16, 18, 23, 25, 27, 29, 31,
			16, 18, 23, 25, 27, 29, 31, 33,
			18, 23, 25, 27, 29, 31, 33, 36,
			23, 25, 27, 29, 31, 33, 36, 38,
			25, 27, 29, 31, 33, 36, 38, 40,
			27, 29, 31, 33, 36, 38, 40, 42,
		},
		{
			9, 13, 15, 17, 19, 21, 22, 24,
			13, 13, 17, 19, 21, 22, 24, 25,
			15, 17, 19, 21, 22, 24, 25, 27,
			17, 19, 21, 22, 24, 25, 27, 28,
			19, 21, 22, 24, 25, 27, 28, 30,
			21, 22, 24, 25, 27, 28, 30, 32,
			22, 24, 25, 27, 28, 30, 32, 33,
			24, 25, 27, 28, 30, 32, 33, 35,
		},
	}
)

// AVRational ...
type AVRational struct {
	Num int
	Den int
}

// SPS ...
/**
 * Sequence parameter set
 */
type SPS struct {
	SPSID                       uint `json:"sps_id"`
	ProfileIdc                  int  `json:"profile_idc"`
	LevelIdc                    int  `json:"level_idc"`
	ChromaFormatIdc             int  `json:"chroma_format_idc"`
	TransformBypass             int  `json:"transform_bypass"`
	Log2MaxFrameNum             int  `json:"log2_max_frame_num"`
	PocType                     int  `json:"poc_type"`
	Log2MaxPocLsb               int  `json:"logx_max_poc_lsb"`
	DeltaPicOrderAlwaysZeroFlag int  `json:"delta_pic_order_always_zero_flag"`
	OffsetForNonRefPic          int  `json:"offset_for_non_ref_pic"`
	OffsetForTopToBottomField   int  `json:"offset_for_top_to_bottom_field"`
	PocCycleLength              int  `json:"poc_cycle_length"`
	RefFrameCount               int  `json:"ref_frame_count"`
	GapsInFrameNumAllowedFlag   int  `json:"gaps_in_frame_num_allowed_flag"`
	MbWidth                     int  `json:"mb_width"`
	MbHeight                    int  `json:"mb_height"`
	FrameMbsOnlyFlag            int  `json:"frame_mbs_only_flag"`
	MbAff                       int  `json:"mb_aff"`
	Direct8x8InterfaceFlag      int  `json:"direct_8x8_interface_flag"`
	Crop                        int  `json:"crop"`

	/* those 4 are already in luma samples */
	VuiParametersPresentFlag     int `json:"vui_parameters_present_flag"`
	Sar                          AVRational
	VideoSignalTypePresentFlag   int    `json:"video_signal_type_present_flag"`
	FullRange                    int    `json:"full_range"`
	ColourDescriptionPresentFlag int    `json:"colour_description_present_flag"`
	ColorPrimaries               uint   `json:"color_primaries"`
	ColorTrc                     uint   `json:"color_trc"`
	ColorSpace                   uint   `json:"color_space"`
	TimingInfoPresentFlag        int    `json:"timing_info_present_flag"`
	NumUtilsInTick               uint32 `json:"num_utils_in_tick"`
	TimeScale                    uint32 `json:"time_scale"`
	FixedFrameRateFlag           int    `json:"fixed_frame_rate_flag"`
	OffsetForRefFrame            [256]int
	BitStreamRestrictionFlag     int `json:"bit_stream_restriction_flag"`
	NumReorderFrames             int `json:"num_reorder_frames"`
	ScalingMatrixPresent         int `json:"scaling_matrix_present"`
	ScalingMatrix4               [6][16]uint8
	ScalingMatrix8               [6][64]uint8
	NalHrdParametersPresentFlag  int `json:"nal_hrd_parameters_present_flag"`
	VclHrdParametersPresentFlag  int `json:"vcl_hrd_parameters_present_flag"`
	PicStructPresentFlag         int `json:"pic_struct_present_flag"`
	TimeOffsetLength             int `json:"time_offset_length"`
	CpbCnt                       int `json:"cpb_cnt"`
	InitialCpbRemovalDelayLength int `json:"initial_cpb_removal_delay_length"`
	CpbRemovalDelayLength        int `json:"cpb_removal_delay_length"`
	DpbOutputDelayLength         int `json:"dpb_output_delay_length"`
	BitDepthLuma                 int `json:"bit_depth_luma"`
	BitDepthChorma               int `json:"bit_depth_chorma"`
	ResidualColorTransformFlag   int `json:"residual_color_transform_flag"`
	ConstraintSetFlags           int `json:"constraint_set_flags"`
}

// PrintSPS ...
// func (s *SPS) PrintSPS() {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	for k := 0; k < t.NumField(); k++ {
// 		fmt.Printf("%s -> %v\n", t.Field(k).Name, v.Field(k).Interface())
// 	}
// }
