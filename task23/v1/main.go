package main

import "fmt"

/*
Удалить i-ый элемент из слайса.

Быстрый способ, выполняется за постоянное время.
Минус: меняется порядок элементов.
*/

func deleteElement(array []string, i int) []string {
	// копируем последний элемент в индекс i.
	array[i] = array[len(array)-1]
	// записываем нулевое значение в последний элемент
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
