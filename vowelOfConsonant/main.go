package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

/*
	func main() {
		args := os.Args

		if len(args) != 2 || len(args[1]) != 1 {
			fmt.Println("Give me a letter")
			return
		}

		// I didn't use a short-if here because, it's already
		// hard to read. Do not make it harder.

		s := args[1]
		if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
			fmt.Printf("%q is a vowel.\n", s)
		} else if s == "w" || s == "y" {
			fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
		} else {
			fmt.Printf("%q is a consonant.\n", s)
		}
	}
*/
/*
func main() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if strings.IndexAny(s, "aeiou") != -1 {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}

	// Notice that:
	//
	// I didn't use IndexAny for the else if above.
	//
	// It's because, calling a function is a costly operation.
	// And, this way, the code is simpler.
}
*/
/*func main() {
	if len(os.Args) == 1 || len(os.Args[1]) > 1 {
		fmt.Println("Give me a letter")
	} else if os.Args[1] == "a" || os.Args[1] == "e" || os.Args[1] == "i" || os.Args[1] == "o" || os.Args[1] == "u" {
		fmt.Printf("%q is a vowel.\n", os.Args[1])
	} else if os.Args[1] == "y" || os.Args[1] == "w" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", os.Args[1])
	} else {
		fmt.Printf("%q is a consonant.\n", os.Args[1])
	}
}
*/
func main() {
	a1 := os.Args[1]
	if len(os.Args) == 1 || len(a1) > 1 {
		fmt.Println("Give me a letter")
	} else if strings.ContainsAny(a1, "aeiou") {
		fmt.Printf("%q is a vowel.\n", a1)
	} else if a1 == "y" || a1 == "w" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", a1)
	} else {
		fmt.Printf("%q is a consonant.\n", a1)
	}
}
