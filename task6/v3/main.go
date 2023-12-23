package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Используем механизм закрытия канала. В функции проверяем по второму возвращаемому аргументу ok, закрыт ли канал.
Если канал закрыт, горутина останавливается.
*/

var wg sync.WaitGroup

func do(data chan int) {
	defer wg.Done()
	fmt.Println("Goroutine running...")
	for {
		_, ok := <-data
		if !ok {
			fmt.Println("Goroutine closing...")
			return
		}

	}
}

func main() {
	data := make(chan int)
	go do(data)
	wg.Add(1)

	for i := 0; i < 3; i++ {
		data <- i
		time.Sleep(3 * time.Second)
	}
	close(data)
	wg.Wait()
}
