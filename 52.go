package main

var x, y, z = "demo", true, "foo"

func main() {
	var q, r = 789, false
	r, s := true, ""
	r = y
	x = s
	_, _ = r, q
}