package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Используем канал. В функции ожидаем получение значения в канале.
*/

var wg sync.WaitGroup

func do(quit chan bool) {
	defer wg.Done()
	fmt.Println("Goroutine running...")
	<-quit
	fmt.Println("Goroutine closing...")
}

func main() {
	quit := make(chan bool)
	go do(quit)
	wg.Add(1)

	time.Sleep(5 * time.Second)
	quit <- true
	wg.Wait()
}
