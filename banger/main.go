package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	banger := strings.Repeat("!", l) + msg + strings.Repeat("!", l)

	fmt.Println(strings.ToUpper(banger))

}
