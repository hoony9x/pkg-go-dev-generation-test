package main

import (
	"fmt"

	"github.com/hoony9x/pkg-go-dev-generation-test/pkg/sample"
)

func main() {
	s := sample.NewSample("test")
	fmt.Println(s.GetName())

	s.SetName("test2")
	fmt.Println(s.GetName())
}
