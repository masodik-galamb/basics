package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1]
	feet, err := strconv.ParseFloat(args, 64)
	meters := feet * 0.3048

	if err != nil {
		fmt.Printf("Error: '%s' is not a number\n", args)
		return
	}
	fmt.Printf("%g feet is %g meters\n", feet, meters)
}
