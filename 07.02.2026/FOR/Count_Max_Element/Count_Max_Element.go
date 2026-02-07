package main

import (
	"fmt"
	"slices"
)

func main() {
	var num int
	var mas []int

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		} else {
			mas = append(mas, num)
		}
	}

	max := slices.Max(mas)
	count := 0

	for i := 0; i < len(mas); i++ {
		if mas[i] == max {
			count++
		}
	}
	fmt.Print(count)
}
