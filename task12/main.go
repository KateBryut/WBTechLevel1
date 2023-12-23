package main

import (
	"fmt"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	group := make(map[string]string)
	for i, val := range array {
		//проверяем, добавлен ли уже элемент в множество по ключу мапы
		if _, ok := group[val]; !ok {
			group[array[i]] = array[i]
		}
	}
	fmt.Println(group)
}
