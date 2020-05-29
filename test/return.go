package test

import "fmt"

func ReturnTest() {
	g := &gop{
		idx: 0,
	}
	g.write()
	g.write()
	g.write()
	fmt.Println("g.idx:", g.idx)
}

type gop struct {
	idx        int
	prev, next *gop
}

func (g *gop) write() {
	ng := &gop{
		idx:  g.idx + 1,
		prev: g,
	}
	g.next = ng
	g = ng
}
