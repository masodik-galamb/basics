package main

import (
	"fmt"
	"path"
)

//version := 0 // bad idea, use below instead

//var version string

func main() {

	//score := 0 // DON'T
	//var score int // already score = 0

	/*var(
		// related:
		video string

		// closely related:
		duration int
		current int
	)
	*/

	//var width. heith = 100, 50 // DON'T
	//width, heith := 100, 50

	//DON'T
	//width = 50 // assign 50 to width
	//color := "red" // new var color
	//INSTEAD
	//width, color := 50, "red" // same result

	_, file := path.Split("test/index.html")

	//fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)

}
