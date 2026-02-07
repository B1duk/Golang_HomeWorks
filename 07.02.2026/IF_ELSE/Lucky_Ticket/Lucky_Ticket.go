package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	half := num % 1000
	num = num / 1000

	fmt.Println(num, half)

	halfSumm := 0
	numSumm := 0

	for half > 0 {
		halfSumm += half % 10
		half /= 10
	}
	for num > 0 {
		numSumm += num % 10
		num /= 10
	}
	if halfSumm == numSumm {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
