package main

import (
	"time"
)

/*
Реализовать собственную функцию sleep.

Используем бесконечный цикл.
*/

func sleep(second time.Duration) {
	start := time.Now()
	for {
		curr := time.Now()
		diff := curr.Sub(start) / time.Nanosecond
		if diff >= second {
			break
		}
	}
}

func main() {
	sleep(10)
}
