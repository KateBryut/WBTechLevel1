package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.

Реализован пример на основе адаптера для европейской вилки для американской розетки
*/

// USPlug — интерфейс для американских вилок и розеток.
type USPlug interface {
	USPlugIn()
}

// Socket — американская розетка.
type Socket struct {
}

func (s *Socket) Plug(plug USPlug) {
	fmt.Println("Пробуем что-то вставить в американскую розетку.")
	plug.USPlugIn()
}

// Plug — американская вилка.
type Plug struct {
}

func (p *Plug) USPlugIn() {
	fmt.Println("Американская вилка вставлена.")
}

// EuroPlug — евровилка, не поддерживающая интерфейс USPlug.
type EuroPlug struct {
}

func (p *EuroPlug) EuroPlugIn() {
	fmt.Println("Евровилка вставлена.")
}

// Adapter — адаптер евровилки для американской розетки.
type Adapter struct {
	Euro *EuroPlug
}

// USPlugIn — адаптер, поддерживающий USPlug-интерфейс.
func (a *Adapter) USPlugIn() {
	a.Euro.EuroPlugIn()
}

func main() {
	socket := &Socket{}
	socket.Plug(&Plug{})

	euroPlug := &EuroPlug{}
	adapter := &Adapter{
		Euro: euroPlug,
	}

	socket.Plug(adapter)
}
