package main

const y = 789
var x int = 123

func main() {
	println("Outer x:", x)

	var x = true
	println("Inner x (bool):", x)
	{
		x, y := x, y
		println("Inner Block x (bool):", x)
		println("Inner Block y (const):", y)

		x, z := !x, y/10
		println("Inner Block x (negation of bool):", x)
		println("Inner Block z (y/10):", z)

		y /= 100
	}

	println("After Inner Block x (bool):", x)
}