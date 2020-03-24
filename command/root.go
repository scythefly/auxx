package command

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auxx",
	Short: "auxx auxx auxx",
}

func init() {
	rootCmd.AddCommand(
		newLeetcodeCommand(),
		newTestCommand(),
	)
}

func Execute() {
	rootCmd.Execute()
}
