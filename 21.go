package main

import "math/rand"

func StateRandonNumbers(n int) (int, int){
	var a, b int
	for i := 0; i< n; i ++ {
		if rand.Intn(MaxRnd) < MaxRnd/2{
			a = a + 1
		}else {
			b ++
		}
	}
	return a, b
}