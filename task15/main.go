package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/

var justString string

func someFunc() {
	// в памяти хранится большой массив
	v := createHugeString(1 << 10)
	justString = v[:100]
	// при создании нового массива не происходит копирования значений массива в новой области памяти
	// происходит лишь копирование ссылки на кусок массива (первые 100 элементов), остальные будут висеть в памяти

	justString1 := make([]byte, 0, 100)
	justString2 := make([]byte, 0, 100)
	//решение: использование функции copy
	copy(justString1, v)
	//явно задать длину переменной justString, дабавть в нее первые 100 элементов
	justString2 = append(justString2, v[:100]...)
}

func main() {
	someFunc()
}
