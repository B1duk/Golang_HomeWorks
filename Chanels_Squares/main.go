package main

import (
	"fmt"
)

// ПЕРВАЯ ГОРУТИНА ГЕНЕРИРУЕТ ЧИСЛА ОТ 1 ДО N И ПИШЕТ ИХ В КАНАЛ
// ВТОРАЯ ГОРУТИНА ЧИТАЕТ ЧИСЛА, ВОЗВОДИТ В КВАДРАТ И ПИШЕТ В ДРУГОЙ КАНАЛ
// Main СОБИРАЕТ КВАДРАТЫ В СРЕЗ И ВЫВОДИТ
func main() {
	var n int
	fmt.Scan(&n)

	ch := make(chan int, n)

	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()

	ch2 := make(chan int, len(ch))
	go func() {

		for i := range ch {
			ch2 <- i * i
		}

		close(ch2)
	}()

	var result []int
	for val := range ch2 {
		result = append(result, val)
	}
	fmt.Println(result)
}
