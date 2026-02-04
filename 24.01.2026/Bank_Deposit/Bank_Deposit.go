package main

import "fmt"

func main() {
	var x, p, y int
	var years int
	fmt.Scanln(&x, &p, &y)
	for x <= y {
		x = x + x*p/100
		years++
	}
	fmt.Print(years)
}
