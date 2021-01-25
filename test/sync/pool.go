package sync

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/atomic"
)

func newPoolCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool",
		Short: "Run pool examples",
		RunE:  runPool,
	}

	return cmd
}

var (
	// frameBuf = []byte("Get selects an arbitrary item fm the Pool, removes it from the Pool, and returns it to the caller. Get may choose to ignore the pool and treat it as empty. Callers should not assume any relation between values passed to Put and the values returned by Get.")
	frameBuf = make([]byte, 1000*1024)
	_allocs  atomic.Uint32
)

func runPool(_ *cobra.Command, _ []string) error {
	fmt.Println(">>> pid: ", os.Getpid())
	debug.SetGCPercent(50)
	for i := 0; i < len(frameBuf); i++ {
		frameBuf[i] = 97 + byte(i%24)
	}
	// var f *frame
	for i := 0; i < 1000; i++ {
		// fmt.Println("==============> new frame", i)
		if i%60 == 0 {
			fmt.Printf("[%d] =========> allocs: %d\n", i, _allocs.Load())
		}
		// if f == nil {
		// 	f = newFrame(frameBuf)
		// } else {
		// 	f.next = newFrame(frameBuf)
		// 	f = f.next
		// }
		newFrame(frameBuf)
		// runtime.SetFinalizer(f, releaseFrame)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("======> done")
	time.Sleep(3 * time.Second)
	fmt.Println("go gc")
	runtime.GC()
	time.Sleep(3 * time.Second)
	runtime.GC()
	time.Sleep(3 * time.Second)
	runtime.GC()
	time.Sleep(3 * time.Second)
	runtime.GC()
	time.Sleep(3 * time.Second)
	runtime.GC()
	select {}
	// return nil
}

var bufPool = sync.Pool{
	New: func() interface{} {
		_allocs.Inc()
		return &cBuffer{
			buf: make([]byte, 1024*512),
		}
	},
}

type cBuffer struct {
	buf  []byte
	next *cBuffer
}

type frame struct {
	data *cBuffer
	p    *cBuffer

	next *frame
}

func releaseFrame(f *frame) {
	// fmt.Println(">>> releaseFrame")
	for f.data != nil {
		p := f.data
		f.data = f.data.next
		p.next = nil
		bufPool.Put(p)
	}
}

func newFrame(b []byte) *frame {
	data, _ := bufPool.Get().(*cBuffer)
	f := &frame{
		data: data,
	}
	f.p = f.data
	offset := copy(f.data.buf, b)
	for offset < len(b) {
		f.p.next, _ = bufPool.Get().(*cBuffer)
		f.p = f.p.next
		n := copy(f.p.buf, b[offset:])
		offset += n
	}

	runtime.SetFinalizer(f, releaseFrame)
	return f
}
