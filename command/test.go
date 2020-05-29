package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"auxx/test"
)

func newTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "test",
		Short: `Run test examples
	-- ring
	-- buffer
	-- cond
	-- cond1
	-- error
	-- ctx
	-- common-pool
	-- conn
	-- chan
	-- xml
	-- defer
	-- return
	-- url
	-- pointer
	-- timer
	-- img
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "ring":
					test.RingBuffer()
				case "buffer":
					test.Buffer()
				case "cond":
					test.CondTest()
				case "cond1":
					test.Cond1Test()
				case "error":
					test.ErrorTest()
				case "ctx":
					test.CtxTest()
				case "common-pool":
					test.CommonPoolTest()
				case "conn":
					test.ConnTest()
				case "chan":
					var cs int = 1500
					if len(args) > 1 {
						cs, _ = strconv.Atoi(args[1])
					}
					test.ChanTest(cs)
				case "xml":
					test.XmlTest()
				case "defer":
					test.DeferTest()
				case "return":
					test.ReturnTest()
				case "memory":
					test.MemoryTest()
				case "url":
					test.URLTest()
				case "pointer":
					test.PointerTest()
				case "timer":
					test.TimerTest()
				case "img":
					test.ImageTest()
				default:
					fmt.Println("----  unknown test command  ----")
				}
			}
		},
	}

	return cmd
}
