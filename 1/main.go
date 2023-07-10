package main

import "fmt"

func newHuman(name, surname string, age int) *Human {
	return &Human{
		name:    name,
		surname: surname,
		age:     age,
	}
} //создаем экземпляр структуры Human и возвращаем указатель на нее

type Human struct {
	name    string
	surname string
	age     int
} //композиция используется в Go со структурами(объект состоит из более мелких объектов)

func (h *Human) speak() {
	fmt.Printf("%s %s: hello world! I am %d years old", h.name, h.surname, h.age)
} // метод структуры Human, который будет встроен в структуру Action

type Action struct {
	*Human //тип Human встроен в Action
}

func main() {
	person := newHuman("test", "testov", 32) // создаем объет
	action := Action{
		person, //встраиваем объект, с которым будет проводиться работа
	}

	action.speak()
	//как видим, мы можем использовать метод(speak) типа Human прямо, не зная путь к нему.
	//стоит избегать возможное столкновение(коллизию) названий в методах.
}
