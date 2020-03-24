package test

import (
	"container/ring"
	"fmt"
	"reflect"
	"time"

	"github.com/scythefly/orb"
	"golang.org/x/sync/errgroup"
)

type frame struct {
	idx   int
	value int
}

type client struct {
	id  int
	cur int
	in  chan struct{}
	r   *ring.Ring
}

type server struct {
	cur       int
	rr        *ring.Ring
	in        chan struct{}
	clients   orb.Set
	keyFrameR *ring.Ring
}

var (
	g errgroup.Group
)

func newServer() *server {
	s := &server{
		rr:      ring.New(12),
		in:      make(chan struct{}),
		clients: orb.NewSet(),
	}
	s.keyFrameR = s.rr
	return s
}

func (s *server) start() {
	go s.broadcast()
	s.setValue()
}

func (s *server) broadcast() {
	for {
		<-s.in
		s.clients.Each(func(i interface{}) bool {
			if c, ok := i.(*client); ok {
				select {
				case c.in <- struct{}{}:
				default:
				}
			}
			return false
		})
	}
}

func (s *server) setValue() {
	ticker := time.NewTicker(500 * time.Millisecond)
	var cnt int
	var f *frame
	var ok bool
	for cnt < 200 {
		<-ticker.C
		if f, ok = s.rr.Value.(*frame); ok {
			// ??? type of f is nil interface
			if f == nil {
				fmt.Println(">>>>>>> nil nil nil nil")
			} else {
				fmt.Println(">>>>>>> ", f.idx, f.value)
			}
		}
		s.rr.Value = &frame{
			idx:   cnt,
			value: cnt % 12,
		}
		if cnt%6 == 0 {
			s.keyFrameR = s.rr
		}
		s.cur = cnt
		s.rr = s.rr.Next()
		cnt++
		fmt.Println("--------------------  in  --------------------", cnt)
		s.in <- struct{}{}
	}
}

func (s *server) newClient(id int) {
	c := &client{
		id: id,
		in: make(chan struct{}),
		r:  s.keyFrameR,
	}
	s.clients.Add(c)
	//if v, ok := c.r.Value.(*frame); ok {
	//	fmt.Println("client:", c.id, "handle frame[", v.idx, v.value, "]")
	//	time.Sleep(400 * time.Millisecond)
	//}
	for {
		<-c.in
		for c.cur < s.cur {
			v, _ := c.r.Value.(*frame)
			// ??? type of v is nil *frame
			if v == nil {
				fmt.Println("- - - - - - - - - - - value is nil - - - - - - - - - - - - - -")
				t := reflect.TypeOf(c.r.Value)
				if t == nil {
					fmt.Println("- - - - - - nil interface - - - - - - -")
				} else {
					fmt.Println("- - - - - -", t.String())
				}
				break
			}
			if c.cur+1 != v.idx {
				c.r = s.keyFrameR
				v, _ = c.r.Value.(*frame)
			}
			c.cur = v.idx
			fmt.Println("client:", c.id, "handle frame[", v.idx, v.value, "]")
			c.r = c.r.Next()
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func RingBuffer() {
	s := newServer()
	g.Go(func() error {
		s.start()
		return nil
	})
	go func() {
		for cnt := 0; cnt < 10; cnt++ {
			go s.newClient(cnt)
			time.Sleep(time.Second)
		}
	}()

	g.Wait()
}
