package command

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/spf13/cobra"

	"auxx/sample"
	"auxx/version"
)

var (
	plugins = []string{
		"leetcode.so",
		"ui.so",
		"utility.so",
	}

	rootCmd = &cobra.Command{
		Use:          "auxx",
		Short:        "auxx auxx auxx",
		SilenceUsage: true,
		// SilenceErrors: true,
	}
)

func init() {
	rootCmd.AddCommand(
		sample.NewTestCommand(),
		newUpdateCommand(),
		version.NewCommand(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}

func Load() error {
	for idx := range plugins {
		realpath := filepath.Join(os.Args[0], "../", plugins[idx])
		fmt.Println("load plugin:", realpath)
		plug, err := plugin.Open(realpath)
		if err != nil {
			return err
		}
		symbol, err := plug.Lookup("Commander")
		if err != nil {
			return err
		}
		if cmd, ok := symbol.(*cobra.Command); ok {
			rootCmd.AddCommand(cmd)
		} else {
			return fmt.Errorf("Symbol in %s is not a commander", realpath)
		}
	}
	return nil
}
