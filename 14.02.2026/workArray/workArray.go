package main

import "fmt"

func main() {

	var workArray [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	var a1, b1, a2, b2, a3, b3 int

	fmt.Scanln(&a1, &b1)
	fmt.Scanln(&a2, &b2)
	fmt.Scanln(&a3, &b3)
	var res string
	if a1+a2+a3+b1+b2+b3 <= 54 {
		res = "ok"
	} else {
		res = "not ok"
	}
	workArray[a1], workArray[b1] = workArray[b1], workArray[a1]
	workArray[a2], workArray[b2] = workArray[b2], workArray[a2]
	workArray[a3], workArray[b3] = workArray[b3], workArray[a3]

	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
	fmt.Printf(res)

}
