package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	loginErr = "Access denied to %q\n"
	passErr  = "Invalid password %q\n"
	accessOK = "Access granted to %q\n"

	user = "an4kein"
	pass = "1337"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != user {
		fmt.Printf(loginErr, u)
	} else if p != pass {
		fmt.Printf(passErr, u)
	} else {
		fmt.Printf(accessOK, u)
	}

}
