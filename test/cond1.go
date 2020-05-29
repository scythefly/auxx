package test

import (
	"container/ring"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var (
	gCond1 errgroup.Group
)

func Cond1Test() {
	p := newProducer()
	gCond1.Go(func() error {
		p.produce()
		return nil
	})
	gCond1.Go(func() error {
		testReader(p)
		return nil
	})
	gCond1.Wait()
}

type packet struct {
	idx  int
	data []byte
}

type packetReader interface {
	ReadPacket() (*packet, error)
	ID() int
}

type producer struct {
	r, lr *ring.Ring
	idx   int

	cond *sync.Cond
	once sync.Once
}

func newProducer() *producer {
	return &producer{
		r:    ring.New(128),
		cond: sync.NewCond(&sync.RWMutex{}),
	}
}

func (p *producer) produce() {
	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			if p.idx%100 == 0 {
				fmt.Printf(">>> write packet, idx: %d\n", p.idx+1)
			}
			p.r.Value = &packet{
				idx:  p.idx + 1,
				data: []byte("xxxxxxxxxxxxxxxxxx default string xxxxxxxxxxxxxxxxxx"),
			}
			p.lr = p.r
			p.r = p.r.Next()
			p.idx++
			// p.cond.L.Lock()
			p.cond.Broadcast()
			// p.cond.L.Unlock()
		}
	}
}

type reader struct {
	id   int
	r    *ring.Ring
	p    *producer
	idx  int
	cond *sync.Cond
}

func (p *producer) NewReader(id int) packetReader {
	r := &reader{
		id:   id,
		p:    p,
		cond: p.cond,
	}
	if p.lr == nil {
		r.r = p.r
	} else {
		r.r = p.lr
	}
	// r.cond.L.Lock()
	return r
}

func (r *reader) ReadPacket() (*packet, error) {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()
	if r.idx+10 < r.p.idx {
		r.r = r.p.lr
		r.idx = r.p.idx
	}
	if r.idx >= r.p.idx {
		if r.id == 999 {
			// fmt.Printf("reader[%d] wait...\n", r.id)
		}
		r.cond.Wait()
		// time.Sleep(10 * time.Millisecond)
		if r.id == 999 {
			// fmt.Printf("reader[%d] wait finished...\n", r.id)
		}
	}

	if r.idx < r.p.idx {
		pkt, _ := r.r.Value.(*packet)
		r.idx = pkt.idx
		r.r = r.r.Next()
		return pkt, nil
	}
	return nil, errors.New("!!!!!!!!!!!!!!!")
}

func (r *reader) ID() int {
	return r.id
}

func testReader(p *producer) {
	for i := 0; i < 20000; i++ {
		id := i
		rr := p.NewReader(id)
		go func(r packetReader) {
			for {
				// _, err := r.ReadPacket()
				pkt, err := r.ReadPacket()
				if err != nil {
					// fmt.Println(err)
					// return
					continue
				}
				if r.ID() == 999 && pkt.idx%100 == 0 {
					fmt.Printf("reader %d get packet %d...\n", r.ID(), pkt.idx)
				}
				// time.Sleep(2 * time.Second)
			}
		}(rr)
	}
}
