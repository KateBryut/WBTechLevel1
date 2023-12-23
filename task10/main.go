package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

Создаем мапу, где ключ - это целое число с шагом 10, в значение помещаем массив чисел.
*/

func main() {
	m := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float32)
	for i := range m {
		group := int(m[i]/10) * 10
		if groups[group] == nil {
			groups[group] = make([]float32, 0, 10)
			groups[group] = append(groups[group], m[i])
		} else {
			groups[group] = append(groups[group], m[i])
		}
	}
	fmt.Println(groups)
}
