package utility

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultLogString = "%<{ats_client_begin}cqh>&%<{ats_client_first_read}cqh>&%<{ats_client_read_done}cqh>&%<{ats_client_header_done}cqh>&%<{ats_cache_read_begin}cqh>&%<{ats_cache_read_end}cqh>&%<{ats_dns_lookup_begin}cqh>&%<{ats_dns_lookup_end}cqh>&%<{ats_cache_write_begin}cqh>&%<{ats_cache_write_end}cqh>&%<{ats_server_first_connect}cqh>&%<{ats_server_connect}cqh>&%<{ats_server_connect_end}cqh>&%<{ats_server_begin_write}cqh>&%<{ats_server_first_read}cqh>&%<{ats_server_header_done}cqh>&%<{ats_client_first_write}cqh>&%<{ats_client_close}cqh>&%<{ats_server_close}cqh"
)

var (
	defaultHooks []string
)

func newTraceCommand() *cobra.Command {
	defaultHooks = strings.Split(defaultLogString, "&")
	cmd := &cobra.Command{
		Use:   "trace",
		Short: "parse trace log",
		Run: func(cmd *cobra.Command, args []string) {
			parseTrace(args)
		},
	}

	return cmd
}

func parseTrace(args []string) {
	if len(args) < 1 {
		return
	}
	var traceString string
	var hooks []string
	traceString = args[0]
	if len(args) > 1 {
		hooks = strings.Split(args[1], "&")
	} else {
		hooks = defaultHooks
	}
	tms := strings.Split(traceString, "&")

	if len(tms) == len(hooks) {
		for idx, v := range hooks {
			fmt.Printf("%s\t%s\n", v, tms[idx])
		}
	}
}
