package test

import (
	"bufio"
	"io"
	"net"
	"os"
	"time"

	orb "github.com/deckarep/golang-set"
	"golang.org/x/sync/errgroup"
)

var (
	gConn errgroup.Group
	bufs  = orb.NewSet()
)

func ConnTest() {
	gConn.Go(func() error {
		return startServer()
	})

	gConn.Go(func() error {
		put()
		return nil
	})

	gConn.Wait()
}

func startServer() error {
	var err error
	l, err := net.Listen("tcp", "localhost:51909")
	if err != nil {
		return err
	}
	for {
		var c net.Conn
		c, err = l.Accept()
		if err != nil {
			break
		}
		m := make(chan []byte)
		bufs.Add(m)
		go handle(c, m)
	}
	return err
}

func put() {
	ticker := time.NewTicker(40 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			bufs.Each(func(v interface{}) bool {
				m, _ := v.(chan []byte)
				m <- []byte(defaultString)
				return false
			})
		}
	}
}

func handle(c net.Conn, m chan []byte) {
	defer bufs.Remove(m)
	r := bufio.NewReader(c)
	w := bufio.NewWriter(c)
	buf := make([]byte, 4096)
	go r.Read(buf)
	for {
		select {
		case b := <-m:
			w.Write(b)
		}
	}
}

func ttttt(c net.Conn) {
	file, _ := os.Open("xxx")
	defer file.Close()
	w := bufio.NewWriter(c)

	r := bufio.NewReader(file)
	io.Copy(w, r)
}
