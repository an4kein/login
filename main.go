package main

import (
	"fmt"
	"os"
)

const usage = `
Usage: [username] [password]
`

var (
	user = "anakein"
	pass = "1337"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usage)
		return
	}
	username, password := os.Args[1], os.Args[2]
	if username == user && password == pass {
		fmt.Printf("Access granted to  \"%s\n", username)
	} else if username == user && password == pass {
		fmt.Printf("Access denied to  \"%s\"\n", username)
	} else if username == user && !(password == pass) {
		fmt.Printf("Invalid password to  \"%s\"\n", username)
	} else if username == user && password == pass {
		fmt.Printf("Invalid password to \"%s\"\n", username)
	} else {
		fmt.Printf("Access denied to \"%s\"\n", username)
	}
}
