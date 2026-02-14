package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	mas := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&mas[i])
	}
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Print(mas[i], " ")
		}
	}
}
