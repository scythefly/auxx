package h265

/* hevc nal type */
const (
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

	HEVC_NAL_TRAIL_N    = 0
	HEVC_NAL_TRAIL_R    = 1
	HEVC_NAL_TSA_N      = 2
	HEVC_NAL_TSA_R      = 3
	HEVC_NAL_STSA_N     = 4
	HEVC_NAL_STSA_R     = 5
	HEVC_NAL_RADL_N     = 6
	HEVC_NAL_RADL_R     = 7
	HEVC_NAL_RASL_N     = 8
	HEVC_NAL_RASL_R     = 9
	HEVC_NAL_VCL_N10    = 10
	HEVC_NAL_VCL_R11    = 11
	HEVC_NAL_VCL_N12    = 12
	HEVC_NAL_VCL_R13    = 13
	HEVC_NAL_VCL_N14    = 14
	HEVC_NAL_VCL_R15    = 15
	HEVC_NAL_BLA_W_LP   = 16
	HEVC_NAL_BLA_W_RADL = 17
	HEVC_NAL_BLA_N_LP   = 18
	HEVC_NAL_IDR_W_RADL = 19
	HEVC_NAL_IDR_N_LP   = 20
	HEVC_NAL_CRA_NUT    = 21
	HEVC_NAL_IRAP_VCL22 = 22
	HEVC_NAL_IRAP_VCL23 = 23
	HEVC_NAL_RSV_VCL24  = 24
	HEVC_NAL_RSV_VCL25  = 25
	HEVC_NAL_RSV_VCL26  = 26
	HEVC_NAL_RSV_VCL27  = 27
	HEVC_NAL_RSV_VCL28  = 28
	HEVC_NAL_RSV_VCL29  = 29
	HEVC_NAL_RSV_VCL30  = 30
	HEVC_NAL_RSV_VCL31  = 31
	HEVC_NAL_VPS        = 32
	HEVC_NAL_SPS        = 33
	HEVC_NAL_PPS        = 34
	HEVC_NAL_AUD        = 35
	HEVC_NAL_EOS_NUT    = 36
	HEVC_NAL_EOB_NUT    = 37
	HEVC_NAL_FD_NUT     = 38
	HEVC_NAL_SEI_PREFIX = 39
	HEVC_NAL_SEI_SUFFIX = 40
	HEVC_NAL_RSV_NVCL41 = 41
	HEVC_NAL_RSV_NVCL42 = 42
	HEVC_NAL_RSV_NVCL43 = 43
	HEVC_NAL_RSV_NVCL44 = 44
	HEVC_NAL_RSV_NVCL45 = 45
	HEVC_NAL_RSV_NVCL46 = 46
	HEVC_NAL_RSV_NVCL47 = 47
	HEVC_NAL_UNSPEC48   = 48
	HEVC_NAL_UNSPEC49   = 49
	HEVC_NAL_UNSPEC50   = 50
	HEVC_NAL_UNSPEC51   = 51
	HEVC_NAL_UNSPEC52   = 52
	HEVC_NAL_UNSPEC53   = 53
	HEVC_NAL_UNSPEC54   = 54
	HEVC_NAL_UNSPEC55   = 55
	HEVC_NAL_UNSPEC56   = 56
	HEVC_NAL_UNSPEC57   = 57
	HEVC_NAL_UNSPEC58   = 58
	HEVC_NAL_UNSPEC59   = 59
	HEVC_NAL_UNSPEC60   = 6
	HEVC_NAL_UNSPEC61   = 61
	HEVC_NAL_UNSPEC62   = 62
	HEVC_NAL_UNSPEC63   = 63

	HEVC_SLICE_B = 0
	HEVC_SLICE_P = 1
	HEVC_SLICE_I = 2

	HEVC_MAX_LAYERS                  = 63
	HEVC_MAX_SUB_LAYERS              = 7
	HEVC_MAX_LAYER_SETS              = 1024
	HEVC_MAX_VPS_COUNT               = 16
	HEVC_MAX_SPS_COUNT               = 16
	HEVC_MAX_PPS_COUNT               = 64
	HEVC_MAX_DPB_SIZE                = 16
	HEVC_MAX_REFS                    = HEVC_MAX_DPB_SIZE
	HEVC_MAX_SHORT_TERM_REF_PIC_SETS = 64
	HEVC_MAX_LONG_TERM_REF_PICS      = 3
	HEVC_MIN_LOG2_CTB_SIZE           = 4
	HEVC_MAX_LOG2_CTB_SIZE           = 6
	HEVC_MAX_CPB_CNT                 = 32
	HEVC_MAX_LUMA_PS                 = 35651584
	HEVC_MAX_WIDTH                   = 16888
	HEVC_MAX_HEIGHT                  = 16888
	HEVC_MAX_TILE_ROWS               = 22
	HEVC_MAX_TILE_COLUMNS            = 20
	HEVC_MAX_SLICE_SEGMENTS          = 600
	HEVC_MAX_ENTRY_POINT_OFFSETS     = HEVC_MAX_TILE_COLUMNS * 135
)

// ProfileTierLevel profile_tier_level
type ProfileTierLevel struct {
	GeneralProfileSpace                 uint `json:"general_profile_space"`
	GeneralTierFlag                     uint `json:"general_tier_flag"`
	GeneralProfileIdc                   uint `json:"general_profile_idc"`
	GeneralProfileCompatibilityFlag     [32]uint8
	GeneralProgressiveSourceFlag        uint `json:"general_progressive_source_flag"`
	GeneralInterlacedSourceFlag         uint `json:"general_interlaced_source_flag"`
	GeneralNonPackedConstraintFlag      uint `json:"general_non_packed_constraint_flag"`
	GeneralFrameOnlyConstraintFlag      uint `json:"general_frame_only_constraint_flag"`
	GeneralMax12bitConstraintFlag       uint `json:"general_max_12bit_constraint_flag"`
	GeneralMax10bitConstraintFlag       uint `json:"general_max_10bit_constraint_flag"`
	GeneralMax8bitConstraintFlag        uint `json:"general_max_8bit_constraint_flag"`
	GeneralMax422chromaConstraintFlag   uint `json:"general_max_422chroma_constraint_flag"`
	GeneralMax420chromaConstraintFlag   uint `json:"general_max_420chroma_constraint_flag"`
	GeneralMaxMonochromeConstraintFlag  uint `json:"general_max_monochrome_constraint_flag"`
	GeneralIntraConstraintFlag          uint `json:"general_intra_constraint_flag"`
	GeneralOnePictureOnlyConstraintFlag uint `json:"general_one_picture_only_constraint_flag"`
	GeneralLowerBitRateConstraintFlag   uint `json:"general_lower_bit_rate_constraint_flag"`
	GeneralMax14bitConstraintFlag       uint `json:"general_max_14bit_constraint_flag"`
	GeneralInbldFlag                    uint `json:"general_inbld_flag"`
	GeneralLevelIdc                     uint `json:"general_level_idc"`
	//
	SubLayerProfilePresentFlag           [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerLevelPresentFlag             [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerProfileSpace                 [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerTierFlag                     [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerProfileIdc                   [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerProfileCompatibilityFlag     [HEVC_MAX_SUB_LAYERS][32]uint8
	SubLayerProgressiveSourceFlag        [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerInterlacedSourceFlag         [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerNonPackedConstraintFlag      [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerFrameOnlyConstraintFlag      [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMax12bitConstraintFlag       [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMax10bitConstraintFlag       [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMax8bitConstraintFlag        [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMax422chromaConstraintFlag   [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMax420chromaConstraintFlag   [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMaxMonochromeConstraintFlag  [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerIntraConstraintFlag          [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerOnePictureOnlyConstraintFlag [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerLowerBitRateConstraintFlag   [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerMax14bitConstraintFlag       [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerInbldFlag                    [HEVC_MAX_SUB_LAYERS]uint8
	SubLayerLevelIdc                     [HEVC_MAX_SUB_LAYERS]uint8
}

// RawScalingList  H265RawScalingList ...
type RawScalingList struct {
	PredModeFlag      [4][6]uint8
	PredMatrixIDDelta [4][6]uint8
	DcCoefMinus8      [4][6]int16
	DeltaCoeff        [4][6][64]int8
}

// RawSTRefPicSet ...
type RawSTRefPicSet struct {
	InterRefPicSetPredictionFlag uint8                 `json:"inter_ref_pic_set_prediction_flag"`
	DeltaIdxMinus1               uint8                 `json:"delta_idx_minus1"`
	DeltaRpsSign                 uint8                 `json:"delta_rps_sign"`
	AbsDeltaRpsMinus1            uint16                `json:"abs_delta_rps_minus1"`
	UsedByCurrPicFlag            [HEVC_MAX_REFS]uint8  `json:"used_by_curr_pic_flag"`
	UseDeltaFlag                 [HEVC_MAX_REFS]uint8  `json:"use_delta_flag"`
	NumNegativePics              uint8                 `json:"num_negative_pics"`
	NumPositivePics              uint8                 `json:"num_positive_pics"`
	DeltaPocS0Minus1             [HEVC_MAX_REFS]uint16 `json:"delta_poc_s0_minus1"`
	UsedByCurrPicS0Flag          [HEVC_MAX_REFS]uint8  `json:"used_by_curr_pic_s0_flag"`
	DeltaPocS1Minus1             [HEVC_MAX_REFS]uint16 `json:"delta_poc_s1_minus1"`
	UsedByCurrPicS1Flag          [HEVC_MAX_REFS]uint8  `json:"used_by_curr_pic_s1_flag"`
}

// SPS ...
type SPS struct {
	// nal_unit_header
	NalUnitType                          uint `json:"nal_unit_type"`
	SPSVideoParameterSetID               uint `json:"sps_video_parameter_set_id"`
	SPSMaxSubLayersMinus1                uint `json:"sps_max_sub_layers_minus1"`
	SPSTemporalIDNestingFlag             uint `json:"sps_temporal_id_nesting_flag"`
	PTL                                  ProfileTierLevel
	SPSSeqParameterSetID                 uint `json:"sps_seq_parameter_set_id"`
	ChromaFormatIdc                      uint `json:"chroma_format_idc"`
	SeparateColourPlaneFlag              uint `json:"separate_colour_plane_flag"`
	PicWidthInLumaSamples                uint `json:"pic_width_in_luma_samples"`
	PicHeightInLumaSamples               uint `json:"pic_height_in_luma_samples"`
	ConformanceWindowFlag                uint `json:"conformance_window_flag"`
	ConfWinLeftOffset                    uint `json:"conf_win_left_offset"`
	ConfWinRightOffset                   uint `json:"conf_win_right_offset"`
	ConfWinTopOffset                     uint `json:"conf_win_top_offset"`
	ConfWinBottomOffset                  uint `json:"conf_win_bottom_offset"`
	BitDepthLumaMinus8                   uint `json:"bit_depth_luma_minus8"`
	BitDepthChromaMinus8                 uint `json:"bit_depth_chroma_minus8"`
	Log2MaxPicOrderCntLsbMinus4          uint `json:"log2_max_pic_order_cnt_lsb_minus4"`
	SPSSubLayerOrderingInfoPresentFlag   uint `json:"sps_sub_layer_ordering_info_present_flag"`
	SPSMaxDecPicBufferingMinus1          [HEVC_MAX_SUB_LAYERS]uint8
	SPSMaxNumReorderPics                 [HEVC_MAX_SUB_LAYERS]uint8
	SPSMaxLatencyIncreasePlus1           [HEVC_MAX_SUB_LAYERS]uint8
	Log2MinLumaCodingBlockSizeMinus3     uint `json:"log2_min_luma_coding_block_size_minus3"`
	Log2DiffMaxMinLumaCodingBlockSize    uint `json:"log2_diff_max_min_luma_coding_block_size"`
	Log2MinLumaTransformBlockSizeMinus2  uint `json:"log2_min_luma_transform_block_size_minus2"`
	Log2DiffMaxMinLumaTransformBlockSize uint `json:"log2_diff_max_min_luma_transform_block_size"`
	MaxTransformHierarchyDepthInter      uint `json:"max_transform_hierarchy_depth_inter"`
	MaxTransformHierarchyDepthIntra      uint `json:"max_transform_hierarchy_depth_intra"`
	ScalingListEnabledFlag               uint `json:"scaling_list_enabled_flag"`
	SPSScalingListDataPresentFlag        uint `json:"sps_scaling_list_data_present_flag"`
	ScalingList                          RawScalingList
	AmpEnabledFlag                       uint `json:"amp_enabled_flag"`
	SampleAdaptiveOffsetEnabledFlag      uint `json:"sample_adaptive_offset_enabled_flag"`
	PcmEnabledFlag                       uint `json:"pcm_enabled_flag"`
	PcmSampleBitDepthLumaMinus1          uint `json:"pcm_sample_bit_depth_luma_minus1"`
	PcmSampleBitDepthChromaMinus1        uint `json:"pcm_sample_bit_depth_chroma_minus1"`
	Log2MinPcmLumaCodingBlockSizeMinus3  uint `json:"log2_min_pcm_luma_coding_block_size_minus3"`
	Log2DiffMaxMinPcmLumaCodingBlockSize uint `json:"log2_diff_max_min_pcm_luma_coding_block_size"`
	PcmLoopFilterDisabledFlag            uint `json:"pcm_loop_filter_disabled_flag"`
	NumShortTermRefPicSets               uint `json:"num_short_term_ref_pic_sets"`
	NumDeltaPocs                         [HEVC_MAX_SHORT_TERM_REF_PIC_SETS]uint
	STRefPicSet                          [HEVC_MAX_SHORT_TERM_REF_PIC_SETS]RawSTRefPicSet
	LongTermRefPicsPresentFlag           uint `json:"long_term_ref_pics_present_flag"`
	NumLongTermRefPicsSPS                uint `json:"num_long_term_ref_pics_sps"`
	LtRefPicPocLsbSPS                    [HEVC_MAX_LONG_TERM_REF_PICS]uint16
	UsedByCurrPicLtSPSFlag               [HEVC_MAX_LONG_TERM_REF_PICS]uint8
	SPSTemporalMvpEnabledFlag            uint `json:"sps_temporal_mvp_enabled_flag"`
	StrongIntraSmoothingEnabledFlag      uint `json:"strong_intra_smoothing_enabled_flag"`
	VuiParametersPresentFlag             uint `json:"vui_parameters_present_flag"`
	// vui
	SPSExtensionPresentFlag    uint `json:"sps_extension_present_flag"`
	SPSRangeExtensionFlag      uint `json:"sps_range_extension_flag"`
	SPSMultilayerExtensionFlag uint `json:"sps_multilayer_extension_flag"`
	SPS3dExtensionFlag         uint `json:"sps_3d_extension_flag"`
	SPSSccExtensionFlag        uint `json:"sps_scc_extension_flag"`
	SPSExtension4bits          uint `json:"sps_extension_4bits"`
	// extension_data
	TransformSkipRotationEnabledFlag          uint `json:"transform_skip_rotation_enabled_flag"`
	TransformSkipContextEnabledFlag           uint `json:"transform_skip_context_enabled_flag"`
	ImplicitRdpcmEnabledFlag                  uint `json:"implicit_rdpcm_enabled_flag"`
	ExplicitRdpcmEnabledFlag                  uint `json:"explicit_rdpcm_enabled_flag"`
	ExtendedPrecisionProcessingFlag           uint `json:"extended_precision_processing_flag"`
	IntraSmoothingDisabledFlag                uint `json:"intra_smoothing_disabled_flag"`
	HighPrecisionOffsetsEnabledFlag           uint `json:"high_precision_offsets_enabled_flag"`
	PersistentRiceAdaptationEnabledFlag       uint `json:"persistent_rice_adaptation_enabled_flag"`
	CabacBypassAlignmentEnabledFlag           uint `json:"cabac_bypass_alignment_enabled_flag"`
	SPSCurrPicRefEnabledFlag                  uint `json:"sps_curr_pic_ref_enabled_flag"`
	PaletteModeEnabledFlag                    uint `json:"palette_mode_enabled_flag"`
	PaletteMaxSize                            uint `json:"palette_max_size"`
	DeltaPaletteMaxPredictorSize              uint `json:"delta_palette_max_predictor_size"`
	SPSPalettePredictorInitializerPresentFlag uint `json:"sps_palette_predictor_initializer_present_flag"`
	SPSNumPalettePredictorInitializerMinus1   uint `json:"sps_num_palette_predictor_initializer_minus1"`
	SPSPalettePredictorInitializers           [3][128]uint16
	MotionVectorResolutionControlIdc          uint `json:"motion_vector_resolution_control_idc"`
	IntraBoundaryFilteringDisableFlag         uint `json:"intra_boundary_filtering_disable_flag"`
}
