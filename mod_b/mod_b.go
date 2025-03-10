package mod_b

import (
	"fmt"

	"github.com/pkg/errors"
)

func PrintB() error {
	fmt.Println("This is package B")

	return bar()
}

func foo() error {
	return errors.New("something went wrong")
}

func bar() error {
	return foo()
}
