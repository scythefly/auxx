package types_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
import "github.com/kakami/pkg/types"

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_")

func Test_Consistent(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cc := types.NewConsistent()
	cc.NumberOfReplicas = 1237

	var (
		// ip1 = "61.156.196.160"
		// ip2 = "223.111.250.52"
		// ip3 = "61.147.236.210"
		// ip4 = "42.240.152.10"
		// ip5 = "42.240.152.60"
		// ip1 = fmt.Sprintf("%c%c%c%c", 42, 20, 152, 34)
		// ip2 = fmt.Sprintf("%c%c%c%c", 242, 240, 12, 35)
		// ip3 = fmt.Sprintf("%c%c%c%c", 2, 24, 152, 136)
		// ip4 = fmt.Sprintf("%c%c%c%c", 42, 240, 152, 195)
		// ip5 = fmt.Sprintf("%c%c%c%c", 42, 240, 152, 196)
		// ip1 = "42.240.152.34"
		// ip2 = "42.240.152.35"


		ip1 = fmt.Sprintf("%d",42 << 24 + 240 << 16 + 152 << 8 + 34)
		ip2 = fmt.Sprintf("%d",42 << 24 + 240 << 16 + 152 << 8 + 35)
		ip3 = fmt.Sprintf("%d",42 << 24 + 240 << 16 + 152 << 8 + 36)
		ip4 = fmt.Sprintf("%d",42 << 24 + 240 << 16 + 152 << 8 + 195)
		ip5 = fmt.Sprintf("%d",42 << 24 + 240 << 16 + 152 << 8 + 196)
	)

	cc.Add(ip1)
	cc.Add(ip2)
	cc.Add(ip3)
	cc.Add(ip4)
	cc.Add(ip5)

	var a1, a2, a3, a4, a5 int

	for i := 0; i < 1000; i++ {
		s, _ := cc.Get(randomString(127))
		switch s {
		case ip1:
			a1++
		case ip2:
			a2++
		case ip3:
			a3++
		case ip4:
			a4++
		case ip5:
			a5++
		default:
			t.Errorf("===================")
		}
	}

	fmt.Println(ip1, a1, "\n", ip2, a2, "\n", ip3, a3, "\n", ip4, a4, "\n", ip5, a5)
}

func randomString(length int) string {
	var out []byte
	for i := 0; i < length; i++ {
		out = append(out, chars[rand.Int()%len(chars)])
	}
	return string(out)
}
