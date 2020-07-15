package test

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func ErrorTest() {
	fmt.Printf("%+v\n", errors.Wrap(io.EOF, "err test io.EOF"))

	fmt.Println(errors.WithMessage(io.EOF, "err test io.EOF").Error())
}
