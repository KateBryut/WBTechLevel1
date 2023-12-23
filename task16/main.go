package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func partition(arr []int, low, high int) ([]int, int) {
	//берем последний элемент массива в качестве опорной точки (последний - чтобы не передвигать в конец)
	pivot := arr[high]
	//i - стартовый индекс
	i := low
	for j := low; j < high; j++ {
		//все что меньше pivot переносим влево
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			fmt.Println(arr)
		}
	}
	// перемещаем pivot в i
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		//рекурсивно повторяем операцию по обе стороны от переменной p
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	fmt.Println(quickSortStart([]int{5, 6, 7, 2, 1, 8}))
}
