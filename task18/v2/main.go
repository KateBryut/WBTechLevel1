package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.

С использованием atomic.
*/

type Counter struct {
	value int32
}

func (c *Counter) inc() {
	atomic.AddInt32(&c.value, 1)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			counter.inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print(counter.value)
}
