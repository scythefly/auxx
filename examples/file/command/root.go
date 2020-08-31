package command

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "file",
		RunE:         runE,
		SilenceUsage: true,
	}
	confPath string
)

func init() {
	rootCmd.AddCommand(
		newUpdateCommand(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}

func runE(_ *cobra.Command, _ []string) error {
	ff, err := os.OpenFile("/mfs-mount/testtt/01.m3u8", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		n, err := ff.WriteString(fmt.Sprintf("01010124.ts?sequence=0000000004&starttime=0000000000004&endtime=0000000000004&startsize=00000000000000000004&endsize=0000000000000000000%d\n", i))
		fmt.Println(n, err)
		if err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println(ff.Close())
	return nil
}
