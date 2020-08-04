package update

import (
	"net/http"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Run Update Examples",
		RunE:  updateRun,
	}

	return cmd
}

func updateRun(_ *cobra.Command, _ []string) error {
	return doUpdate("http://localhost:8081/build/auxx")
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}
	return err
}
