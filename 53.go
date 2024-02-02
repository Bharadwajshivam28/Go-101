package main

import "fmt"

func main () {
	const a = -1.23
	var x = int32(a)
	var y int32 = int32(a)

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	const k int16 = 255

	var f = uint8(k + 1)
	var h = uint(k) + 1

	fmt.Println("f:", f)
	fmt.Println("h:", h)
}