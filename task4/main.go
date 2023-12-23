package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.

Выбран способ завершения работы при помощи контекста.
Родительский контекст передается в воркеры, получающие данные.
При нажании Ctrl+C выполняется родительский контекст, при его выполнении завершаются воркеры и паблишер.
*/

var wg sync.WaitGroup

func listen(ctx context.Context, dataChannel chan int, id int) {
	defer wg.Done()
	for {
		select {
		//отслеживаем выполнение родительского контекста
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "Quit worker '%d'.\n", id)
			return
		case <-dataChannel:
			fmt.Fprintf(os.Stdout, "Worker '%d' receive message '%d'.\n", id, <-dataChannel)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func post(ctx context.Context, dataChannel chan int) {
	defer wg.Done()
	for {
		select {
		//отслеживаем выполнение родительского контекста
		case <-ctx.Done():
			fmt.Println("Quit publisher")
			return
		default:
			time.Sleep(400 * time.Millisecond)
			dataChannel <- rand.Int()
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataChannel := make(chan int)
	signalChannel := make(chan os.Signal, 1)
	//создаем подписку на Ctrl+C
	signal.Notify(signalChannel, os.Interrupt)

	var n int
	fmt.Print("Enter num of workers: ")
	fmt.Scan(&n)
	//запускаем n воркеров
	for i := 0; i < n; i++ {
		go listen(ctx, dataChannel, i)
		wg.Add(1)
	}

	//запускаем паблишер
	go post(ctx, dataChannel)
	wg.Add(1)

	<-signalChannel
	//при получении сигнала Ctrl+C отменяем контекст
	cancel()
	wg.Wait()
}
