package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Raw Concat
//
//  1. Initialize the name variable
//     by getting input from the command line
//
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
//
// NOTE
//  You should concatenate the name variable in the correct
//  place.
//
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func main() {
	// uncomment the code below
	// name := "and get the name from the command-line"
	name := os.Args[1]

	// replace and concatenate the `name` variable
	// after `hi ` below

	// msg := `hi CONCATENATE-NAME-VARIABLE-HERE!
	// how are you?`

	/*msg := `hi ` + name + `!
	how are you?`

		fmt.Println(msg)
	*/

	msg := []string{`hi `, `!
how are you?`}

	fmt.Println(strings.Join(msg, name))
}
