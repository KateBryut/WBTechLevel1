package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

//родительская структура
type Human struct {
	Name string
}

//дочерняя структура
type Action struct {
	Human
}

func (h Human) Do() {
	fmt.Printf("Do something")
}

func main() {
	a := Action{}
	a.Name = "Name"
	//вызов метода родительской структуры
	a.Do()
}
