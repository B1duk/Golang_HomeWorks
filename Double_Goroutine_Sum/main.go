package main

import (
	"fmt"
)

func SliceSum(nums []int) int {

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("Результат текущего канала: ", sum)
	return sum

}

func main() {

	// Условие: Есть срез чисел. Нужно посчитать их сумму,
	// но вычисление разбить на две горутины
	// (первая половина и вторая половина среза),
	// а результаты собрать через канал.

	var n int
	fmt.Scan(&n)

	ch := make(chan int, 2)

	if n == 0 {
		return
	}

	nums := make([]int, 0)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		nums = append(nums, m)
	}
	mid := len(nums) / 2
	chunk1 := nums[:mid]
	chunk2 := nums[mid:]
	go func() {
		ch <- SliceSum(chunk1)
	}()

	go func() {
		ch <- SliceSum(chunk2)
	}()

	res1 := <-ch
	res2 := <-ch

	fmt.Println(res1 + res2)

}
