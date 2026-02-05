package main

import (
	"fmt"
	"math"
)

func main() {

	//в интеренте написано что a-угол броска надо указывать в радианах
	var v, a float64
	g := 9.8 //ускорение свободного падения

	fmt.Scanln(&v, &a)

	a = (a * math.Pi) / 180

	H := (math.Pow(v, 2) * math.Pow(math.Sin(a), 2)) / (2 * g)
	L := (math.Pow(v, 2) * math.Sin(2*a)) / g
	T := (2 * v * math.Sin(a)) / g
	fmt.Printf("Максимальная высота: %.2f \n Дальность полета: %.2f \n Время полета: %.2f", H, L, T)
}
