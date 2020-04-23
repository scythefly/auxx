package test

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/jolestar/go-commons-pool/v2"
)

func CommonPoolTest() {
	simpleExample()
	customExample()
}

func simpleExample() {
	type myPoolObject struct {
		s string
	}

	v := uint64(0)
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			return &myPoolObject{
					s: strconv.FormatUint(atomic.AddUint64(&v, 1), 10),
				},
				nil
		})

	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)

	obj, err := p.BorrowObject(ctx)
	if err != nil {
		panic(err)
	}

	o := obj.(*myPoolObject)
	fmt.Println(o.s)

	err = p.ReturnObject(ctx, obj)
	if err != nil {
		panic(err)
	}
}

type MyPoolObject struct {
	s string
}

type MyCustomFactory struct {
	v uint64
}

func (f *MyCustomFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	return pool.NewPooledObject(
			&MyPoolObject{
				s: strconv.FormatUint(atomic.AddUint64(&f.v, 1), 10),
			}),
		nil
}

func (f *MyCustomFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	// do destroy
	return nil
}

func (f *MyCustomFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	// do validate
	return true
}

func (f *MyCustomFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do activate
	return nil
}

func (f *MyCustomFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do passivate
	return nil
}

func customExample() {
	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, &MyCustomFactory{})
	p.Config.MaxTotal = 100

	obj1, err := p.BorrowObject(ctx)
	if err != nil {
		panic(err)
	}

	o := obj1.(*MyPoolObject)
	fmt.Println(o.s)

	err = p.ReturnObject(ctx, obj1)
	if err != nil {
		panic(err)
	}

	// Output: 1
}
