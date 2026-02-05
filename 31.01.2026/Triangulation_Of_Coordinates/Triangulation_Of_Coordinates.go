package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3 float64

	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	fmt.Scanln(&x3, &y3)

	a := math.Hypot(x2-x1, y2-y1)
	b := math.Hypot(x3-x2, y3-y2)
	c := math.Hypot(x1-x3, y1-y3)
	p := (a + b + c) / 2
	S := math.Sqrt(p * (p - 2) * (p - b) * (p - c))

	fmt.Printf("%.2f", S)
}
