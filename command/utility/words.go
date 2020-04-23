package utility

import "github.com/spf13/cobra"

type wordsOption struct {
	unitFrom int
	uintTo   int
	unit     int
	amount   int
}

func newFetchWordsCommand() *cobra.Command {
	opt := &wordsOption{
	}
	cmd := &cobra.Command{
		Use:   "words",
		Short: "get some words from db",
		Run: func(cmd *cobra.Command, args []string) {
			fetchWords(opt, args)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opt.unit, "unit-from", "s", -1, "words after this unit(included)")
	flags.IntVarP(&opt.unit, "unit-to", "e", -1, "words before this unit(-1 means the end)")
	flags.IntVarP(&opt.amount, "amount", "n", 10, "the amount of words to output")

	return cmd
}

func fetchWords(opt *wordsOption, args []string) {
	openDB()
}
