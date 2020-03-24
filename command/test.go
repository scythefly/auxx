package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"auxx/test"
)

func newTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Run test examples",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "ring":
					test.RingBuffer()
				case "buffer":
					test.Buffer()
				case "cond":
					test.CondTest()
				default:
					fmt.Println("----  unknown test command  ----")
				}
			}
		},
	}

	return cmd
}
