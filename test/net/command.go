package net

import (
	"github.com/spf13/cobra"

	"auxx/test/net/http"
	"auxx/test/net/http/fasthttp"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "net",
		Short: "Run http examples",
	}

	cmd.AddCommand(
		http.NewCommand(),
		fasthttp.NewCommand(),
		newConnCommand(),
	)

	return cmd
}
