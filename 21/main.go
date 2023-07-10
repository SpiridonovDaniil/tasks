package main

import (
	"fmt"
)

// Вариант 1
type contract interface { //интерфейс, с которым должна работать наша система.
	work() string
}

type oldService struct { //пользовательский тип, который система должна адаптировать под себя.
}

func (o *oldService) work() string { //метод, который мы бы хотели использовать.
	return "result old service"
}

func newAdapter(oldService *oldService) contract { //конструктор адаптера.
	return &adapter{oldService}
}

type adapter struct { //адаптер, реализующий нужный нам интерфейс.
	*oldService
}

func main() {
	old := oldService{}
	adapt := newAdapter(&old)
	fmt.Println(adapt.work())
}

//TODO разобраться
