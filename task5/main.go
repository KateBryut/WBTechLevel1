package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func listen(ctx context.Context, dataChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		//отслеживаем выполнение родительского контекста по таймауту
		case <-ctx.Done():
			return
		case <-dataChannel:
			fmt.Fprintf(os.Stdout, "Worker '%d' receive message.\n", <-dataChannel)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func post(ctx context.Context, dataChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		//отслеживаем выполнение родительского контекста по таймауту
		case <-ctx.Done():
			return
		default:
			time.Sleep(400 * time.Millisecond)
			dataChannel <- rand.Int()
		}
	}
}

func main() {
	var n string
	fmt.Print("Enter timeout: ")
	fmt.Scan(&n)
	timeout, _ := time.ParseDuration(n)
	//завершение программы реализуем через контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	dataChannel := make(chan int)
	defer close(dataChannel)

	wg.Add(2)
	go post(ctx, dataChannel, &wg)
	go listen(ctx, dataChannel, &wg)
	wg.Wait()
}
