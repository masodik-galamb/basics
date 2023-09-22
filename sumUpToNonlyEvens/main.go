package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: Provide two different non-zero numeric arguments")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil ||
		min <= 0 || max <= 0 || min >= max {
		fmt.Println("Wrong numbers")
		return
	}
	var sum int

	for i := min; i <= max; i++ {

		switch {
		case i%2 != 0:
			continue
		default:
			sum += i
			fmt.Print(i)
			if i > min && i < max {
				fmt.Print(" + ")
			}
		}
	}
	fmt.Printf(" = %d\n", sum)
}
