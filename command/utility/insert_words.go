package utility

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newInsertWordsCommand() *cobra.Command {
	opt := &wordsOption{
	}
	cmd := &cobra.Command{
		Use:   "add-words",
		Short: "insert some words to db",
		Run: func(cmd *cobra.Command, args []string) {
			insertWords(opt, args)
		},
	}
	flags := cmd.Flags()
	flags.IntVar(&opt.unit, "unit", 0, "specified the unit of words")

	return cmd
}

func insertWords(opt *wordsOption, args []string) {
	openDB()
	var err error
	sqlString := `create table if not exists tb_junior_one (
	id INTEGER PRIMARY KEY,
	unit INTEGER,
	word TEXT NOT NULL
);`
	if _, err = db.Exec(sqlString); err != nil {
		panic(err)
	}

	if len(args) < 1 {
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		if len(line) > 1 {
			word := line[0 : len(line)-2]
			sqlString = fmt.Sprintf(`insert into tb_junior_one(unit,word) values(%d,'%s');`, opt.unit, word)
			db.Exec(sqlString)
		}
		if err != nil {
			break
		}
	}
}
