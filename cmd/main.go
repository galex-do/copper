package main

import (
	"fmt"

	"github.com/galex-do/copper/pkg/module1"
	"github.com/galex-do/copper/pkg/module2"
)

func main() {
	fmt.Println("Initializing Copper Application")

	module1.DoSomething()
	module2.DoSomethingElse()
}
