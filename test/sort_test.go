package test_test

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

type node struct {
	idx int32
}

type nodenode struct {
	nodes []*node
}

func (nn *nodenode) Swap(i, j int) {
	nn.nodes[i], nn.nodes[j] = nn.nodes[j], nn.nodes[i]
}

func (nn *nodenode) Less(i, j int) bool {
	return nn.nodes[j].idx >= nn.nodes[i].idx
}

func (nn *nodenode) Len() int {
	return len(nn.nodes)
}

func Test_Sort(t *testing.T) {
	fmt.Println(strconv.FormatInt(time.Now().Unix()+60, 16))
	fmt.Println(strconv.FormatInt(time.Now().Unix()+60, 10))
	nn := &nodenode{}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 10; i++ {
		nn.nodes = append(nn.nodes, &node{idx: r.Int31()})
	}
	tt := time.Now()
	// for idx := range nn.nodes {
	// 	fmt.Println(nn.nodes[idx].idx)
	// }
	fmt.Println("=====================================")
	sort.Sort(nn)
	for idx := range nn.nodes {
		fmt.Println(nn.nodes[idx].idx)
	}
	fmt.Println(">> ", time.Now().Sub(tt))
}

func Test_Match(t *testing.T) {
	t.Log(match("aaa.bbb.ccc.ddd.com", "ccc.com"))
	t.Log(match("aaa.bbb.ccc.ddd.com", "ddd.com"))
	t.Log(match("aaa.bbb.ccc.ddd.com", ".ccc.com"))
	t.Log(match("aaa.bbb.ccc.ddd.com", strings.TrimLeft("*.ddd.com", "*")))
	t.Log(match("aaa.bbb.ccc.ddd.com", ".com"))
}

func match(a, b string) bool {
	if a == b {
		return true
	}

	idx := strings.Index(a, ".")
	for idx > 0 {
		a = a[idx:]
		if a == b {
			return true
		}
		a = a[1:]
		idx = strings.Index(a, ".")
	}
	return false
}
