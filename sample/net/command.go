package net

import (
	"github.com/spf13/cobra"

	"auxx/sample/net/http"
	"auxx/sample/net/http/fasthttp"
	"auxx/sample/net/rpc"
	"auxx/sample/net/writev"
	"auxx/sample/net/ws"
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
