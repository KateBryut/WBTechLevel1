package main

import (
	"time"
)

/*
Реализовать собственную функцию sleep.

Используем таймер.
*/

func sleep(second time.Duration) {
	t := time.NewTimer(second * time.Second)
	<-t.C
}

func main() {
	sleep(10)
}
