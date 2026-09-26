package main

import (
	"fmt"
	"sync"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numElements := len(nums)

	if numElements == 0 {
		return
	}
	if n > numElements {
		n = numElements
	}

	chunkSize := numElements / n
	remainder := numElements % n

	var wg sync.WaitGroup

	start := 0
	for i := 0; i < n; i++ {
		end := start + chunkSize
		if i < remainder {
			end++
		}

		chunk := nums[start:end]

		wg.Add(1)
		go func(subset []int) {
			defer wg.Done()
			sum := 0
			for i := 0; i < len(subset); i++ {
				sum += subset[i]
			}
			fmt.Println("Срез: ", subset)
			fmt.Println("Сумма среза: ", sum)
		}(chunk)

		start = end

	}
	wg.Wait()
	fmt.Print("Done")
}
