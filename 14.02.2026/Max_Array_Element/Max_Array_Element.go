package main

import (
	"fmt"
	"math"
)

func main() {
	var mas [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&mas[i])
	}
	max := math.MinInt
	for i := 0; i < 5; i++ {
		if mas[i] > max {
			max = mas[i]
		}
	}
	fmt.Print(max)
}
