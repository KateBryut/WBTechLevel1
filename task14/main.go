package main

import (
	"fmt"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

// getType - определяет тип переменной с использованием конструкции switch
func getType(any interface{}) string {
	switch any.(type) {
	case string:
		return "string"
	case int:
		return "untyped int"
	case int64:
		return "int64"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan"
	default:
		return "unknown type"
	}
}
func main() {
	fmt.Println(getType("test string"))
	fmt.Println(getType(int(12)))
	fmt.Println(getType(int64(15)))
	fmt.Println(getType(true))
	var ch chan interface{}
	fmt.Println(getType(ch))
}
