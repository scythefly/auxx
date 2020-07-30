package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"auxx/test"
	"auxx/test/etcd"
	"auxx/test/kafka"
	"auxx/test/net"
	"auxx/test/sync"
)

func newTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "test",
		Short: `Run test examples
	-- buffer
	-- common-pool
	-- cond
	-- cond1
	-- conn
	-- ctx
	-- defer
	-- error
	-- img
	-- interface
	-- path
	-- plugin
	-- pointer
	-- return
	-- ring
	-- timer
	-- url
	-- user
	-- xml
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "array":
					test.ArrayTest()
				case "ring":
					test.RingBuffer()
				case "buffer":
					test.Buffer()
				case "cond":
					test.CondTest()
				case "cond1":
					test.Cond1Test()
				case "conf":
					test.ConfTest()
				case "error":
					test.ErrorTest()
				case "ctx":
					test.CtxTest()
				case "common-pool":
					test.CommonPoolTest()
				case "conn":
					test.ConnTest()
				case "xml":
					test.XmlTest()
				case "defer":
					test.DeferTest()
					test.ImageTest()
				case "reflect":
					test.ReflectTest()
				case "return":
					test.ReturnTest()
				case "memory":
					test.MemoryTest()
				case "url":
					test.URLTest()
				case "pointer":
					test.PointerTest()
				case "quic":
					test.QuicTest()
				case "schedule":
					test.ScheduleTest()
				case "timer":
					test.TimerTest()
				case "ticker":
					test.TickerTest()
				case "img":
					test.ImageTest()
				case "path":
					test.PathTest()
				case "plugin":
					test.PluginTest()
				case "interface":
					test.InterfaceTest()
				case "user":
					test.UserTest()
				default:
					fmt.Println("----  unknown test command  ----")
				}
			}
		},
	}

	cmd.AddCommand(
		etcd.NewCommand(),
		net.NewCommand(),
		kafka.NewCommand(),
		sync.NewCommand(),
	)
	return cmd
}
