package command

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var updateOption struct {
	addr string
	uri  string
}

func newUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update yourself",
		RunE:  updateRun,
	}

	flags := cmd.Flags()
	flags.StringVarP(&updateOption.addr, "addr", "r", "update.scythefly.top:61910", "remote addr")
	flags.StringVar(&updateOption.uri, "uri", "download/binary/file", "remote uri")

	return cmd
}

func updateRun(_ *cobra.Command, _ []string) error {
	var addr, uri string
	if addr = os.Getenv("UGTP_UPDATE_ADDR"); addr == "" {
		addr = updateOption.addr
	}
	if uri = os.Getenv("UGTP_UPDATE_URI"); uri == "" {
		uri = updateOption.uri
	}
	return doUpdate("http://" + filepath.Join(addr, uri))
}

func doUpdate(url string) error {
	fmt.Println(">> update from ", url)
	resp, err := http.Get(url)
	if err != nil {
		return errors.WithMessagef(err, "fetch %s", url)
	}
	defer resp.Body.Close()
	if err = update.Apply(resp.Body, update.Options{}); err != nil {
		return errors.WithMessage(err, "update local file")
	}
	return err
}
