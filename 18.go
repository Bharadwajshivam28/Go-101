package main

import "math/rand"

const MaxRnd = 16

func StateRandomNumbers(n int) (int, int) {
	var a, b int
	for i := 0; i< n; i ++ {
		if rand.Intn(MaxRnd) < MaxRnd/2 {
			a = a + 1
		}else {
			b ++
		}
	}
	return a, b
}

func main () {
	var num = 100
	x, y := StateRandomNumbers(num)
	print("Result: ", x, " + ", y, " = ", num, "? ")
	print(x + y == num)
}