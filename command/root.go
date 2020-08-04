package command

import (
	"github.com/spf13/cobra"

	"auxx/command/utility"
	"auxx/version"
)

var rootCmd = &cobra.Command{
	Use:   "auxx",
	Short: "auxx auxx auxx",
}

func init() {
	rootCmd.AddCommand(
		newLeetcodeCommand(),
		newTestCommand(),
		newHttpCommand(),
		utility.NewCommand(),
		version.NewCommand(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}
