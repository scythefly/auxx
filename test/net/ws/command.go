package ws

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ws",
		RunE: runWs,
	}
	return cmd
}

func runWs(_ *cobra.Command, _ []string) error {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8082", nil)
	if err != nil {
		return err
	}
	defer c.Close()
	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read loop:", err)
				return
			}
			fmt.Println("read:", string(msg))
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var cnt int
	for {
		select {
		case t := <-ticker.C:
			cnt++
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				fmt.Println("write loop:", err, cnt)
				return err
			}
			if cnt > 10 {
				fmt.Println("write ", cnt, "quit")
				time.Sleep(time.Second)
				return nil
			}
		}
	}
}
