package main

import "fmt"

func main() {

	var num int

	fmt.Scan(&num)

	if num > 0 {
		fmt.Print("Число положительное")
	} else if num < 0 {
		fmt.Print("Число отрицательное")
	} else {
		fmt.Print("Ноль")
	}
}
