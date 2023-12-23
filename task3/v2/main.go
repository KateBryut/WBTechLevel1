package main

import (
	"fmt"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений

Используем каналы
*/

func main() {
	array := [5]int32{2, 4, 6, 8, 10}
	ch := make(chan int32, len(array))
	sum := int32(0)

	for _, val := range array {
		go func(val int32) {
			s := int32(val * val)
			ch <- s
		}(val)
	}

	for i := 0; i < len(array); i++ {
		sum += <-ch
	}
	fmt.Println(sum)
}
