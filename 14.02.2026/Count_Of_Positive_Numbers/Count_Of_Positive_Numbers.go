package main

import "fmt"

func main() {
	var N, count int
	fmt.Scan(&N)
	mas := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&mas[i])
	}
	for i := 0; i < N; i++ {
		if mas[i] > 0 {
			count++
		}
	}
	fmt.Print(count)
}
