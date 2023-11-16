package main

import (
	"backend/core"
	"fmt"
)

func main() {
	fmt.Println("initial server")
	// conn := connector.NewConnector("bookrzn", "book1995", "127.0.0.1:3306", "bookrzn")
	// conn.AddUser("user", "pswd", "Bob", "+7900000", "ex@", "1234567890")
	// conn.FindUserFromToken("0001")
	// conn.CountRows("bookrzn.Users")
	Core := core.NewCore()
	Core.SetHandlers()
}
