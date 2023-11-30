package main

import (
	"backend/core"
	"fmt"
)

func main() {
	fmt.Println("initial server")
	// sender.Sendeer()
	Core := core.NewCore()
	Core.SetHandlers()
}
