package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.

С использованием мьютекса.
*/

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
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
