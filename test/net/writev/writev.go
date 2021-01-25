package writev

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "writev",
		Short: "Run writev examples",
		RunE:  runWritev,
	}
	return cmd
}

func runWritev(_ *cobra.Command, _ []string) error {
	var bv [][]byte
	for i := 0; i < 5; i++ {
		bv = append(bv, []byte("12333"))
	}
	nv := net.Buffers(bv)
	n, err := nv.WriteTo(os.Stdout)
	fmt.Println("writev:", n)

	return err
}
