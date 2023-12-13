package main

import (
	"backend/connector"
	"backend/core"
	"fmt"
)

func main() {
	fmt.Printf("\n[ SERVER ] -- [ INIT ] -- [ %v ]", connector.DateTimeNow())
	Core := core.NewCore()
	Core.SetHandlers()
}
