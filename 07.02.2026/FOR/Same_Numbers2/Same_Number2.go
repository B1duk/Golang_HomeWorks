package main

import (
	"fmt"
	"strconv"
)

func main() {

	var first, second int

	fmt.Scanln(&first, &second)

	firststr := strconv.Itoa(first)
	secondstr := strconv.Itoa(second)

	firstdigits := []int{}
	seconddigits := []int{}
	for _, char := range firststr {
		digit := int(char - '0')
		firstdigits = append(firstdigits, digit)
	}
	for _, char := range secondstr {
		digit := int(char - '0')
		seconddigits = append(seconddigits, digit)
	}
	sameNumbers := []int{}
	for _, dig := range firstdigits {
		for _, dig2 := range seconddigits {
			if dig == dig2 {
				sameNumbers = append(sameNumbers, dig)
			}
		}
	}

	for _, dig := range sameNumbers {
		fmt.Printf("%d, ", dig)
	}
}
