package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 5
	b := 10
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d, b = %d", a, b)
}
