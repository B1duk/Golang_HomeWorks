package main

import (
	"fmt"
)

func main() {
	var n, num int

	var mas []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		mas = append(mas, num)
	}
	summ := 0
	for i := 0; i < n; i++ {
		if mas[i] >= 10 && mas[i] <= 99 && mas[i]%8 == 0 {
			summ += mas[i]
		}
	}
	fmt.Print(summ)

}
