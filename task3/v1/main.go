package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений

Используем atomic
*/

func main() {
	array := [5]int32{2, 4, 6, 8, 10}
	sum := int32(0)
	var wg sync.WaitGroup

	for _, val := range array {
		wg.Add(1)
		go func(val int32) {
			defer wg.Done()
			s := int32(val * val)
			atomic.AddInt32(&sum, s)
		}(val)
	}
	wg.Wait()
	fmt.Println(sum)
}
