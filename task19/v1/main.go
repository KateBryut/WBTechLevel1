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

Используем конкатенацию строк.
*/

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var output string = ""
	array := strings.Split(input, "")
	mid := int(float64(len(array)) / 2)
	length := len(array)
	for i := 0; i < mid; i++ {
		array[i], array[length-i-1] = array[length-i-1], array[i]
	}
	for i := range array {
		output += array[i]
	}
	fmt.Println(output)
}
