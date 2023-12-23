package main

import (
	"fmt"
	"math"
	"strings"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func changeBit(input, num int64, value int) (int64, error) {
	if value > 1 {
		return 0, fmt.Errorf("value shouldn't be more than 1")
	}
	//конвертируем число в биты и разделяем на массив
	bits := strings.Split(fmt.Sprintf("%064b", input), "")
	//заменяем i-ый байт
	bits[num] = fmt.Sprint(value)

	length := len(bits)
	out := 0.00
	//переводим из 2ой системы в 10ую
	for i := length - 1; i >= 0; i-- {
		if bits[i] == "1" {
			out += math.Pow(2, float64(length-i-1))
		}
	}
	return int64(out), nil
}

func main() {
	val1, err := changeBit(1234, 56, 0)
	if err != nil {
		fmt.Printf("error: %x", err.Error())
		return
	}
	fmt.Printf("new value 1 = %d.\n", val1)
	val2, err := changeBit(1234, 5, 1)
	if err != nil {
		fmt.Printf("error: %x", err.Error())
		return
	}
	fmt.Printf("new value 2 = %d.\n", val2)
}
