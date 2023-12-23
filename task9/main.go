package main

import (
	"fmt"
	"os"
	"time"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

// generator - записывает в канал поочередно числа в бесконечном цикле
func generator(chInput chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chInput <- i
		i++
	}
}

// counter - читает данные из канала chInput и записывает в канал chOut квадрат полученного числа
func counter(chInput, chOut chan int) {
	for {
		chOut <- 2 * <-chInput
	}
}

// printer - читает данные из канала chOut и выводит
func printer(chOut chan int) {
	for {
		fmt.Fprintf(os.Stdout, "Received value: %d \n", <-chOut)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generator(ch1)
	go counter(ch1, ch2)
	go printer(ch2)

	time.Sleep(20 * time.Second)
}
