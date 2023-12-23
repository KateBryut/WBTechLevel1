package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/*
Реализовать конкурентную запись данных в map

Используем sync.Map.
*/

func write(m sync.Map, key, value string) {
	fmt.Fprintf(os.Stdout, "Write to map[%s] value %s...\n", key, value)
	m.Store(key, value)
}

func main() {
	m := sync.Map{}
	wg := sync.WaitGroup{}

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			write(m, fmt.Sprintf("%d", rand.Int()), "value1")
			wg.Done()
		}()
	}

	wg.Wait()
}
