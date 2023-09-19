package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

/*
	func main() {
		a := len(os.Args[1:])

		if a <= 1 || a > 2 {
			fmt.Println("Usage: [username] [password]")
		} else if os.Args[1] == "jack" && os.Args[2] == "1888" {
			fmt.Printf("Access granted to %q.\n", os.Args[1])
		} else if (os.Args[1] != "jack" && os.Args[2] != "1888") || (os.Args[1] != "jack" && os.Args[2] == "1888") {
			fmt.Printf("Access denied for %q.\n", os.Args[1])
		} else if os.Args[1] == "jack" && os.Args[2] != "1888" {
			fmt.Printf("Invalid password for %q.\n", os.Args[1])
		}
	}
*/
/*func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}
	u, p := args[1], args[2]
	if u != "jack" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "1888" {
		fmt.Printf("Invalid password for %q.\n", u)
	} else {
		fmt.Printf("Access granted to %q.\n", u)
	}
}
*/
const (
	usage   = "Usage: [username] [password]"
	errUser = "Access denied for %q.\n"
	errPass = "Invalid password for %q.\n"
	sAccess = "Access granted to %q.\n"

	user = "jack"
	pass = "1888"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != user {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPass, u)
	} else {
		fmt.Printf(sAccess, u)
	}
}
