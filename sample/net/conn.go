package net

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func newConnCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "conn",
		Short: "Run conn examples",
		RunE:  connRun,
	}

	return cmd
}

func connRun(_ *cobra.Command, _ []string) error {
	l, err := net.Listen("tcp", ":30001")
	if err != nil {
		return err
	}
	defer l.Close()

	go connClient()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleRequest(conn)
	}
	return nil
}

func handleRequest(conn net.Conn) {
	time.Sleep(5 * time.Second)
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
	conn.Close()
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
}

func connClient() {
	time.Sleep(2 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:30001")
	if err != nil {
		fmt.Println("client conn dial", err)
		return
	}

	// time.Sleep(5 * time.Second)
	fmt.Println("client", conn.RemoteAddr().String())
	fmt.Println("client", conn.LocalAddr().String())
	conn.Close()
	fmt.Println("client", conn.RemoteAddr().String())
	fmt.Println("client", conn.LocalAddr().String())
}
