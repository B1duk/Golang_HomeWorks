package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)
	mas := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&mas[i])
	}
	if N >= 4 {
		fmt.Print(mas[3])
	} else {
		fmt.Print("N<4")
	}

}
