package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»

Используем конкатеннацию строк.
*/

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var output string = ""
	input = input[:len(input)-2]
	array := strings.Split(input, " ")
	mid := int(float64(len(array)) / 2)
	length := len(array)
	for i := 0; i < mid; i++ {
		array[i], array[length-i-1] = array[length-i-1], array[i]
	}
	for i := range array {
		output += array[i]
		if i < length {
			output += " "
		}
	}
	fmt.Println(output)
}
