package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.

Для решения задачи использован пакет для работы с большими числами - math/big
*/

func sum(a, b *big.Int) *big.Int {
	var res big.Int
	res.Add(a, b)
	return &res
}

func sub(a, b *big.Int) *big.Int {
	var res big.Int
	res.Sub(a, b)
	return &res
}

func mul(a, b *big.Int) *big.Int {
	var res big.Int
	res.Mul(a, b)
	return &res
}

func div(a, b *big.Int) *big.Int {
	var res big.Int
	res.Quo(a, b)
	return &res
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("24000000000000000000", 10)
	b.SetString("25000000000000000000", 10)

	fmt.Println(sum(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
}
