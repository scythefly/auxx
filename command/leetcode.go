package command

import (
	"strconv"

	"github.com/spf13/cobra"

	"auxx/leetcode"
)

func newLeetcodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leetcode",
		Short: "Do leetcode",
		Run: func(cmd *cobra.Command, args []string) {
			var index int
			if len(args) > 0 {
				index, _ = strconv.Atoi(args[0])
			}
			leetcode.Do(index)
		},
	}

	return cmd
}
