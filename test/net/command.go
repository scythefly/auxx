package net

import (
	"github.com/spf13/cobra"

	"auxx/test/net/http"
	"auxx/test/net/http/fasthttp"
	"auxx/test/net/rpc"
	"auxx/test/net/writev"
	"auxx/test/net/ws"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "net",
		Short: "Run http examples",
	}

	cmd.AddCommand(
		fasthttp.NewCommand(),
		http.NewCommand(),
		rpc.NewCommand(),
		writev.NewCommand(),
		ws.NewCommand(),
		newConnCommand(),
	)

	return cmd
}
