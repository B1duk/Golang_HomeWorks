package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)

	if num/100 == (num%100)/10 && (num%100)/10 == num%10 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
