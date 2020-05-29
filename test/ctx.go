package test

import (
	"context"
	"fmt"
)

type ctxTestContent struct {
	Idx  int
	Desc string
}

type ctxKey int

var ctxContKey ctxKey

func NewContext(ctx context.Context, cont *ctxTestContent) context.Context {
	return context.WithValue(ctx, ctxContKey, cont)
}

func FromContext(ctx context.Context) (*ctxTestContent, bool) {
	c, ok := ctx.Value(ctxContKey).(*ctxTestContent)
	return c, ok
}

func CtxTest() {
	ctx := context.Background()
	cont := &ctxTestContent{
		Idx:  0,
		Desc: "from main routine",
	}
	cctx := NewContext(ctx, cont)
	onPlay(cctx)
	fmt.Printf("CtxTest >> idx: %d, desc: %s\n", cont.Idx, cont.Desc)
}

func onPlay(ctx context.Context) {
	if cont, ok := FromContext(ctx); ok {
		cont.Idx++
		cont.Desc = "from onPlay func"
		fmt.Printf("onPlay >> idx: %d, desc: %s\n", cont.Idx, cont.Desc)
	}
}
