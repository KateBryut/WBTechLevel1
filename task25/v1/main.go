package main

import (
	"context"
	"time"
)

/*
Реализовать собственную функцию sleep.

Используем контекст с таймаутом.
*/

func sleep(second time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*second)
	defer cancel()

	<-ctx.Done()
}

func main() {
	sleep(10)
}
