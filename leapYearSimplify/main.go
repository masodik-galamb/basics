package main

import (
	"fmt"
	"os"
	"strconv"
)

var leap bool

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}
	a := os.Args[1]
	year, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", a)
		return
	}
	if year%400 == 0 || year%4 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
