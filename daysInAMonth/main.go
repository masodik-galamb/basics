package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name.")
		return
	}

	args := os.Args[1]
	month := strings.ToLower(os.Args[1])
	year := time.Now().Year()

	if month == "january" ||
		month == "march" ||
		month == "may" ||
		month == "july" ||
		month == "august" ||
		month == "october" ||
		month == "december" {
		fmt.Printf("%q has 31 days\n", args)
	} else if month == "february" &&
		(year%4 == 0 && (year%100 != 0 || year%400 == 0)) {
		fmt.Printf("%q has 29 days\n", args)
	} else if month == "february" &&
		(year%4 != 0 && (year%100 == 0 || year%400 != 0)) {
		fmt.Printf("%q has 28 days\n", args)
	} else if month == "april" ||
		month == "june" ||
		month == "september" ||
		month == "november" {
		fmt.Printf("%q has 30 days\n", args)
	} else {
		fmt.Printf("%q is not a month\n", args)
	}
}
