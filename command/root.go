package command

import (
	"github.com/spf13/cobra"

	"auxx/command/utility"
	"auxx/version"
)

var rootCmd = &cobra.Command{
	Use:          "auxx",
	Short:        "auxx auxx auxx",
	SilenceUsage: true,
	// SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(
		newLeetcodeCommand(),
		newTestCommand(),
		newHttpCommand(),
		newUpdateCommand(),
		// ui.NewCommand(),
		utility.NewCommand(),
		version.NewCommand(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}
