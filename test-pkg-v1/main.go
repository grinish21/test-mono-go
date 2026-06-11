package main

import (
	"fmt"

	"github.com/grinish21/test-mono-go/test-pkg-v1-v2/who"
)

func main() {
	fmt.Println("I am pkg-v1 only")
	who.WhoAmI()
}

func WhoAmI() string {
	return "I am pkg-v1"
}
