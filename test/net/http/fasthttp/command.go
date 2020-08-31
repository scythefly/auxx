package fasthttp

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fast",
		Short: "Run fasthttp examples",
		RunE:  fastRun,
	}
	return cmd
}

func fastRun(_ *cobra.Command, _ []string) error {
	var err error
	c := &fasthttp.HostClient{
		Addr: "update.scythefly.top:61910",
	}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://download.scythefly.top/download/binary/ugtp")
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	file, err := os.Create("ugtp")
	if err != nil {
		return errors.WithMessage(err, "create file")
	}
	defer file.Close()

	if err = c.DoCopy(req, resp, file); err != nil {
		return errors.WithMessage(err, "download file")
	}

	return nil
}
