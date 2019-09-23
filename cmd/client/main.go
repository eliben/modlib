package main

import (
	"fmt"

	"github.com/eliben/modlib"
	"github.com/eliben/modlib/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", modlib.Config())
	clientlib.Hello()
}
