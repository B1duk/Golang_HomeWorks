package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		square = 2
	)
	weight, height := userInput()
	indexCalculation(weight, height, square)

}

func userInput() (float64, float64) {

	var weight, height float64

	fmt.Scanln(&weight, &height)

	return weight, height
}

func indexCalculation(weight, height, square float64) {
	height = height / 100
	heightSquared := math.Pow(height, square)
	index := weight / heightSquared

	if index < 18.5 {
		fmt.Printf("Индекс: %.2f - недостаточная масса тела\n", index)
	} else if index >= 18.5 && index <= 24.9 {
		fmt.Printf("Индекс: %.2f - нормальная масса тела\n", index)
	} else if index >= 25.0 && index <= 29.9 {
		fmt.Printf("Индекс: %.2f - избыточная масса тела\n", index)
	} else if index >= 30.0 {
		fmt.Printf("Индекс: %.2f - ожирение\n", index)
	}

}
