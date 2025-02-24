package main

import (
	"fmt"
	"go-module/c4f"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world!")
	c4f.Hello()
	fmt.Println(quote.Hello())
}
