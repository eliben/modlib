package main

import (
	"fmt"

	"github.com/eliben/modlib"
	"github.com/eliben/modlib/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", modlib.Config())
	serverlib.Hello()
}
