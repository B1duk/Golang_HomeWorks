package main

import (
	"fmt"
)

func main() {
	/*

		  Задача 5: Расчет аннуитетного платежа
		Рассчитать ежемесячный аннуитетный платеж по формуле:
		A = S * (i * (1 + i)^n) / ((1 + i)^n - 1)
		где S - сумма кредита, i - месячная процентная ставка, n - количество месяцев.

	*/

	var S, i, n float64

	fmt.Scanln(&S, &i, &n)
	result := 1.0
	g := 1 + i
	for j := 0; j < int(n); j++ {
		result *= g
	}

	A := S * (i * result) / (result - 1)

	fmt.Printf("%.2f", A)

}
