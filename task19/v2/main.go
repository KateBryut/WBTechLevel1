package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.

Используем strings.Builder. Более предпочтительный вариант.
*/

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	array := strings.Split(input, "")
	var builder strings.Builder

	for i := len(array) - 1; i >= 0; i-- {
		builder.WriteString(array[i])
	}
	fmt.Print(builder.String())
}
