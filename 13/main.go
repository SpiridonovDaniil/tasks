package main

import "fmt"

func main() {
	a := 3
	b := 5

	c := "yes"
	d := "no"

	a = a + b //вариант 1 арифметика
	b = a - b
	a = a - b

	c, d = d, c //вариант 2 используем плюсы множественного присваивания

	fmt.Printf("%d\n%d\n%s\n%s\n", a, b, c, d)
}
