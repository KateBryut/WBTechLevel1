package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Используем контекст. В функции ожидаем выполнение родительского контекста.
*/

var wg sync.WaitGroup

func do(ctx context.Context) {
	defer wg.Done()
	fmt.Println("Goroutine running...")
	<-ctx.Done()
	fmt.Println("Goroutine closing...")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go do(ctx)
	wg.Add(1)

	wg.Wait()
}
