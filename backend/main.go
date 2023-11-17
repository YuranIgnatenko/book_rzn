package main

import (
	"backend/core"
	"fmt"
)

func main() {
	fmt.Println("initial server")

	Core := core.NewCore()
	Core.SetHandlers()
}
