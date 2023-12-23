package main

import "fmt"

/*
Удалить i-ый элемент из слайса.

Выполняется за линейное время.
Плюс: сохраняется порядок элементов.
*/

func deleteElement(array []string, i int) []string {
	// выполняем сдвиг a[i+1:] влево на один индекс
	copy(array[i:], array[i+1:])
	// удяляем последний элемент
	array[len(array)-1] = ""
	// усекаем срез
	array = array[:len(array)-1]

	return array
}

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	fmt.Println(deleteElement(a, i))
}
