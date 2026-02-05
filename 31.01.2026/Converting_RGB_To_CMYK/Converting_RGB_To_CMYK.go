package main

import (
	"fmt"
	"math"
)

func main() {
	/*

		  Преобразование RGB в CMYK
		Конвертировать значения RGB (0-255) в CMYK (0-1) по формулам:
		R' = R/255, G' = G/255, B' = B/255
		K = 1 - max(R', G', B')
		C = (1 - R' - K) / (1 - K)
		M = (1 - G' - K) / (1 - K)
		Y = (1 - B' - K) / (1 - K)
		При K = 1 все значения CMYK равны 0.

	*/

	var R, G, B float64

	fmt.Scanln(&R, &G, &B)

	R = R / 255
	G = G / 255
	B = B / 255
	if RGBCheck(R, G, B) == true {
		K := 1 - math.Max(math.Max(R, G), B)

		C := (1 - R - K) / (1 - K)
		M := (1 - G - K) / (1 - K)
		Y := (1 - B - K) / (1 - K)

		fmt.Printf("%.2f %.2f %.2f %.2f", C, M, Y, K)
	} else {
		fmt.Println("Возникла ошибка")
	}
}

func RGBCheck(R, G, B float64) bool {
	return R >= 0 && R <= 255 && G >= 0 && G <= 255 && B >= 0 && B <= 255
}
