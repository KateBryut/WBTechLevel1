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

Используем strings.Builder.
*/

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = input[:len(input)-2]
	array := strings.Split(input, " ")
	var builder strings.Builder

	for i := len(array) - 1; i >= 0; i-- {
		builder.WriteString(array[i])
		if i < len(array) {
			builder.WriteString(" ")
		}
	}
	fmt.Print(builder.String())
}
