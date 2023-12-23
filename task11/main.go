package main

import (
	"fmt"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	array1 := []int{5, 8, 24, 1, 148}
	array2 := []int{24, 3, 1, 47, 39, 63}
	//мапа, где ключ - значение из меньшего среза, а значение - пустая структура
	mapArray1 := make(map[int]struct{})
	inter := make([]int, 0)

	//выявляем меньший срез по длине
	if len(array1) > len(array2) {
		array1, array2 = array2, array1
	}

	//заполняем мапу значениями меньшего среза
	for _, v := range array1 {
		mapArray1[v] = struct{}{}
	}

	for _, v := range array2 {
		//ищем значение из большего среза ключ в мапе
		//если нашли, добавляем в пересечение
		if _, ok := mapArray1[v]; ok {
			inter = append(inter, v)
		}
	}
	fmt.Print(inter)
}
