package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, val := range array {
		wg.Add(1)
		//используем вызов анонимной ф-ии в разных горутинах
		go func(val int) {
			fmt.Println(val * val)
			wg.Done()
		}(val)
	}
	//дождемся завершения всх горутин
	wg.Wait()
}
