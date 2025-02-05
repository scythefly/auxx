package lrc

import (
	"bufio"
	"io"

	"auxx/plugins/utility/container/pb"
)

func Decode(r io.Reader) []pb.Pair {
	br := bufio.NewReader(r)
	var err error
	var line string
	var state int
	for {
		line, err = br.ReadString('\n')
		if line != "" {
		}
	}
}
