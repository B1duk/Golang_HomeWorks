package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	summ := 0
	for i := num1; i <= num2; i++ {
		summ += i
	}
	fmt.Print(summ)
}
