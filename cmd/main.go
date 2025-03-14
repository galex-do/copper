package main

import (
    "fmt"
    "copper/pkg/module1"
    "copper/pkg/module2"
)

func main() {
    fmt.Println("Initializing Copper Application")
    
    module1.DoSomething()
    module2.DoSomethingElse()
}