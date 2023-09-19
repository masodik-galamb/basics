package main

import "fmt"

func main() {
	speed := 100 // it's int
	force := 3.3 // it's float64

	//speed = speed * int(force)

	speed = int(float64(speed) * force)

	fmt.Println(speed)
	fmt.Println(force, int(force))

	/*float64(100) = 100.0 -> float64
	100.0 * 3.3 = 330.0 -> float64
	int(330.0) = 330 -> int
	*/
}
