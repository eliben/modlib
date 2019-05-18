package main

import (
	"fmt"

	"github.com/eliben/modlib"
	"github.com/eliben/modlib/otherlib"
)

func main() {
	fmt.Println("From modlib:", modlib.Foo1())
	fmt.Println("From modlib:", modlib.Foo3())

	fmt.Println("From otherlib:", otherlib.Bob1())
}
