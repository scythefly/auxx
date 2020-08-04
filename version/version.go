package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "unknown"
	ChangeLog = "unknown"
	Built     = "unknown"
)

// NewVersionCommand ...
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "show version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`AUXX
    Author: Scythefly
    Version: %s
    ChangeLog: %s
    build with %s, at %s
`, Version, ChangeLog, runtime.Version(), Built)
		},
	}

	return cmd
}
