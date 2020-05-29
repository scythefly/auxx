package command

import (
	"github.com/spf13/cobra"

	"auxx/command/utility"
)

var rootCmd = &cobra.Command{
	Use:   "auxx",
	Short: "auxx auxx auxx",
}

func init() {
	rootCmd.AddCommand(
		newLeetcodeCommand(),
		newTestCommand(),
		utility.NewCommand(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}
