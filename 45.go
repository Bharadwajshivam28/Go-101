package main

import "fmt"

const (
	Monday  = iota
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
	Sunday
)

func main () {
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Tues: %d\n", Tuesday)
	fmt.Printf("wed: %d\n", Wednesday)
	fmt.Printf("Thrus: %d\n", Thrusday)
	fmt.Printf("Fri: %d\n", Friday)
	fmt.Printf("Sat: %d\n", Saturday)
	fmt.Printf("Sun: %d\n", Sunday)
}