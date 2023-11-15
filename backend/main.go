package main

import (
	"backend/connector"
	"fmt"
)

func main() {
	fmt.Println("initial server")
	connector.Connect()
	// Core := core.NewCore()
	// Core.SetHandlers()
}
