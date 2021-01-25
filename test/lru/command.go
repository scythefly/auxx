package lru

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/simplelru"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lru",
		RunE: lruRunE,
	}

	// cmd.AddCommand()
	return cmd
}

func lruRunE(_ *cobra.Command, _ []string) error {
	onEvicted := func(k interface{}, v interface{}) {
		fmt.Println("evited:", k, v)
	}
	l, err := lru.NewLRU(90, onEvicted)
	if err != nil {
		return err
	}
	for i := 0; i < 100; i++ {
		l.Add(i, i)
	}
	fmt.Println(l.Keys())
	for i := 0; i < 5; i++ {
		l.Add(i, i)
	}
	fmt.Println(l.Keys())
	for i := 5; i > 1; i-- {
		l.Add(i, i)
	}
	fmt.Println(l.Keys())
	return nil
}
