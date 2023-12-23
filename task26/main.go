package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func checkUniqueness(input string) bool {
	groups := make(map[string]string)
	//приводим к верхнему регистру
	input = strings.ToLower(input)
	//разделяем строку в массив
	arr := strings.Split(input, "")
	for _, v := range arr {
		//проверяем уникальность по наличию ключа в мапе
		if _, ok := groups[v]; ok {
			return false
		}
		groups[v] = v
	}
	return true
}

func main() {
	fmt.Println(checkUniqueness("abcd"))
	fmt.Println(checkUniqueness("abCdefAaf"))
	fmt.Println(checkUniqueness("aabcd"))
}
