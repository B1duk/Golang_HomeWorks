package main

import "fmt"

func main() {
	var matrix [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	a := matrix[0][0]
	b := matrix[0][1]
	c := matrix[0][2]
	d := matrix[1][0]
	e := matrix[1][1]
	f := matrix[1][2]
	g := matrix[2][0]
	h := matrix[2][1]
	i := matrix[2][2]

	output := (a*e*i + b*f*g + c*d*h) - (c*e*g + a*f*h + b*d*i)

	fmt.Println(output)
}
