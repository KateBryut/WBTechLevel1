package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/*
Реализовать конкурентную запись данных в map

Используем мьютекс.
*/

type Map struct {
	sync.RWMutex
	Data map[string]string
}

func write(m *Map, key, value string) {
	m.Lock()
	defer m.Unlock()
	fmt.Fprintf(os.Stdout, "Write to map[%s] value %s...\n", key, value)
	m.Data[key] = value
}

func main() {
	m := &Map{}
	m.Data = make(map[string]string)

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
