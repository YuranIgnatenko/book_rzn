package main

import (
	"backend/core"
	"fmt"
)

func main() {
	fmt.Println("init server")
	Core := core.NewCore()
	Core.SetHandlers()
}
