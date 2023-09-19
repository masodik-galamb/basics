// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go alien
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 9.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	var richter string

	desc := args[1]

	switch desc {
	case "micro":
		richter = "less than 2.0"
	case "very minor":
		richter = "2 - 2.9"
	case "minor":
		richter = "3 - 3.9"
	case "light":
		richter = "4 - 4.9"
	case "moderate":
		richter = "5 - 5.9"
	case "strong":
		richter = "6 - 6.9"
	case "major":
		richter = "7 - 7.9"
	case "great":
		richter = "8 - 9.9"
	case "massive":
		richter = "10+"
	default:
		richter = "unknown"
	}

	fmt.Printf(
		"%s's richter scale is %s\n",
		desc, richter,
	)
}
*/

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	args := os.Args[1]
	var rich int
	var ter float64

	if args == "micro" {
		ter = 2.0
		fmt.Printf("%s's richter scale is less than %.1f\n", args, ter)
		return
	}
	if args == "massive" {
		rich = 10
		fmt.Printf("%s's richter scale is %d+\n", args, rich)
		return
	}
	switch {
	case args == "very minor":
		rich = 2
		ter = 2.9
	case args == "minor":
		rich = 3
		ter = 3.9
	case args == "light":
		rich = 4
		ter = 4.9
	case args == "moderate":
		rich = 5
		ter = 5.9
	case args == "strong":
		rich = 6
		ter = 6.9
	case args == "major":
		rich = 7
		ter = 7.9
	case args == "great":
		rich = 8
		ter = 9.9
	case args == "massive":
		rich = 8
		ter = 9.9
	default:
		fmt.Printf("%s's richter scale is uknown\n", args)
		return
	}
	fmt.Printf("%s's richter scale is %d - %.1f\n", args, rich, ter)
}