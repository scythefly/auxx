package utility

import (
	"fmt"
	"reflect"
	"time"

	"github.com/spf13/cobra"
)

type State struct {
	// ip
	LocalAddr          string        `json:"local_addr"` // $local_addr
	RemoteAddr         string        // $remote_addr
	Domain             string        // $domain
	App                string        // $app
	Name               string        // $name
	TcURL              string        // $tcUrl
	PArgs              string        // $pargs
	UnixTime           time.Time     // $unix_time
	ConnectionTime     time.Time     // $connection_time
	InBytes            uint64        // $in_bytes
	OutBytes           uint64        // $out_bytes
	Frozen             int           // $frozen
	ReferenceFps       int           // $reference_fps
	MinFps             int           // $min_fpx
	SessionType        string        // $sessiontype - relay_pull/relay_push/publisher/player
	Scheme             string        // $scheme
	Status             int           // $status
	StreamSource       string        // $streamsource - internel
	FirstMetaTime      time.Time     // $firstmeta_time
	UserAgent          string        // $useragent - ua/flashversion
	LogStage           string        // $log_stage - start/end/underway
	OclpStatus         int           // $oclp_status
	FinalizeReason     string        // $finalize_reason
	Stage              string        // $stage
	Init               time.Time     // $init
	HandshakeDone      time.Time     // $handshake_done
	Connect            time.Time     // $connect
	CreateStream       time.Time     // $create_stream
	PTime              time.Time     // $ptime - play/publish command time
	FirstData          time.Time     // $first_data
	FirstAudio         time.Time     // $first_audio
	FirstVideo         time.Time     // $first_video
	CloseStream        time.Time     // $close_stream
	Duration           time.Duration // $duration
	Reconnect          int           // $reconnect
	MinAFps            int           // $min_afps
	AudioCodec         string        // $audio_codec
	VideoCodec         string        // $video_codec
	AudioHeaderTime    time.Time     // $audio_header_time
	VideoHeaderTime    time.Time     // $video_header_time
	HlsPlusOnlineCount int           // $hls_plus_onlinecount

	VideoFrames uint64
	AudioFrames uint64
}

func newDecodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "decode struct variables",
		Run: func(cmd *cobra.Command, args []string) {
			decodeStruct()
		},
	}

	return cmd
}

func decodeStruct() {
	t := reflect.TypeOf(State{})
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		fmt.Println(t.Field(i).Name)
	}
}
