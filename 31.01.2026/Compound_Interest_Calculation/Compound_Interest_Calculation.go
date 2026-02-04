package main

import "fmt"

func main() {
	var P, r, n float64
	fmt.Scanln(&P, &r, &n)
	result := 1.0
	g := (1 + r/100)

	for i := 0.0; i < n; i++ {
		result *= g
	}

	S := P * result
	fmt.Printf("%.2f", S)
}
