package main

import (
	"fmt"
	"strconv"
)

/*
Реализовать бинарный поиск встроенными методами языка.

Функция binarySearch в качестве параметра принимает срез, по которому будет проводиться поиск, а также искомый элемент
Возвращает индекс найденного элемента и true в случае нахождения, false -> в противном сулчае
*/

func main() {
	array := []int{5, 6, 7, 2, 1, 0}
	lookingFor := 7
	var assumption, result = binarySearch(&array, lookingFor)
	if result {
		fmt.Println("Found: " + strconv.Itoa(assumption))
	} else {
		fmt.Println("Nothing found!")
	}
}

func binarySearch(array *[]int, lookingFor int) (int, bool) {
	var mid, assumption int
	//вычисляем минимальный и максимальный индексы
	min := 0
	high := len(*array) - 1

	for min <= high {
		//вычисляем среднее значение
		mid = (min + high) / 2
		assumption = mid
		if assumption == lookingFor {
			return assumption, true
		}
		//определяем с какой стороны решение
		//передвигаем границу диапазона
		if assumption > lookingFor {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return 0, false
}
