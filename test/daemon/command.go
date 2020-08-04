package daemon

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"auxx/test"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "Run daemon examples",
		RunE:  daemonRun,
	}
	return cmd
}

func daemonRun(cmd *cobra.Command, args []string) error {
	var err error
	if os.Getppid() != 1 {
		test.Daemon()
	}

	for i := 0; i < 20; i++ {
		fmt.Println(">>>>>>>>>. test", args)
		time.Sleep(3 * time.Second)
	}
	return err
}
