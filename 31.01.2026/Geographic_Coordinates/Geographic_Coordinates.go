package main

import (
	"fmt"
	"math"
)

func main() {

	/*
	   	Географические координаты
	   По координатам двух точек на сфере (широта/долгота в градусах) вычислить расстояние по формуле гаверсинусов:
	   a = sin²(Δφ/2) + cos(φ1) * cos(φ2) * sin²(Δλ/2)
	   c = 2 * atan2(√a, √(1−a))
	   d = R * c
	   Все тригонометрические функции должны использовать градусы.
	*/

	var la1, lo1, la2, lo2 float64
	R := 6371.0
	fmt.Scanln(&la1, &lo1)
	fmt.Scanln(&la2, &lo2)
	la1 = DegreesToRadians(la1)
	lo1 = DegreesToRadians(lo1)
	la2 = DegreesToRadians(la2)
	lo2 = DegreesToRadians(lo2)

	la3 := la2 - la1
	lo3 := lo2 - lo1
	a := math.Pow(math.Sin(la3/2), 2) + (math.Cos(la1) * math.Cos(la2) * math.Pow(math.Sin(lo3/2), 2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := R * c

	fmt.Printf("%.2f", d)
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
