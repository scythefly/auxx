package http

import (
	"net/http"
	"sync"

	"github.com/spf13/cobra"

	"auxx/test/net/http/gin"
)

var (
	dispatchMux *http.ServeMux
	tString     string
	once        sync.Once
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run http examples",
		RunE:  httpRun,
	}

	cmd.AddCommand(
		newUgtpCommand(),
		gin.NewCommand(),
	)

	return cmd
}

func httpRun(*cobra.Command, []string) error {
	dispatchMux = http.NewServeMux()
	dispatchMux.HandleFunc("/test/11", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("......... /test/1111"))
		once.Do(func() {
			dispatchMux.HandleFunc("/test/222", func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("......../test/2222222222"))
			})
		})
	})
	dispatchMux.HandleFunc("/test/22", handleRedirect)
	return http.ListenAndServe(":8989", dispatchMux)
}
