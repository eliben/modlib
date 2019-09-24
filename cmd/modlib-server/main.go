package main

import (
	"fmt"

	"github.com/eliben/modlib"
	"github.com/eliben/modlib/internal/auth"
	"github.com/eliben/modlib/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", modlib.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
