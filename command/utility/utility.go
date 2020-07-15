package utility

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	g      errgroup.Group
	db     *sql.DB
	dbPath string
	once   sync.Once
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go",
		Short: "go go go",
	}

	cmd.AddCommand(
		newsplitFileCommand(),
		newTraceCommand(),
		newFetchWordsCommand(),
		newInsertWordsCommand(),
		newDecodeCommand(),
	)

	flags := cmd.Flags()
	flags.StringVarP(&dbPath, "db-path", "d", "/Users/iuz/Local/db/auxx.db", "sqlite3 database file path")

	return cmd
}

func openDB() {
	once.Do(func() {
		var err error
		if db, err = sql.Open("sqlite3", dbPath); err != nil {
			panic(err)
		}
	})
}
