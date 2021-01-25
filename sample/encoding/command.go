package encoding

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encoding",
		Short: "run encoding examples",
	}

	cmd.AddCommand(
		newXmlCommand(),
	)
	return cmd
}
