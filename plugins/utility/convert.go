package main

import "github.com/spf13/cobra"

var convertOpt struct {
	input  string
	format string
}

func convertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "convert format",
		RunE:  runConvert,
	}

	flags := cmd.PersistentFlags()
	flags.StringVarP(&convertOpt.input, "input", "i", "", "input file path")
	flags.StringVarP(&convertOpt.format, "format", "f", "srt", "output format")

	return cmd
}

func runConvert(_ *cobra.Command, _ []string) error {
	var err error

	return err
}
