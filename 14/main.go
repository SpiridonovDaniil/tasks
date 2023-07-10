package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 4 //инициализируем тестовые данные.
	str := "test"
	bl := true
	ch := make(chan int)

	fmt.Println("вариант 1") //определяем тип через функцию typeOf пакета reflect.
	fmt.Println(worker(num))
	fmt.Println(worker(str))
	fmt.Println(worker(bl))
	fmt.Println(worker(ch))
	fmt.Println("\nвариант 2") //определяем тип через приведение интерфейса.
	fmt.Println(doType(num))
	fmt.Println(doType(str))
	fmt.Println(doType(bl))
	fmt.Println(doType(ch))
}

func worker(i interface{}) reflect.Type {
	res := reflect.TypeOf(i)
	return res
}

func doType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	}
	return "не поддерживаемый тип"
}
