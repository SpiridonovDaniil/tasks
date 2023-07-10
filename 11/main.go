package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //для получения случайного числа при каждом запуске.
	setA := make([]int, 6)           //инициализируем слайс(первое множество).
	for i, _ := range setA {
		setA[i] = rand.Intn(30) //заполняем первое множество случайными числами от 0 до 30.
	}
	setB := make([]int, 8) //инициализируем слайс(второе множество).
	for i, _ := range setB {
		setB[i] = rand.Intn(30) //заполняем второе множество случайными числами от 0 до 30.
	}
	res, err := search(setA, setB) //ищем пересечения.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res) //выводим значения пересечений в стандартный поток вывода.
}

func search(a, b []int) ([]int, error) {
	m := make(map[int]int)  //инициализируем карту.
	res := make([]int, 0)   //инициализируем слайс.
	for _, val := range a { //итерируемся по первому множеству.
		m[val] += 1 //используем значения множества как ключи, а значения увеличиваем на 1.
	}
	for _, val := range b { //итерируемся по второму множеству.
		m[val] += 1 //используем значения множежества как ключи, а значения увеличиваем на 1.
	} //если ключи совпадут, то значения будут равны 2.
	for key, val := range m { //итерируемся по карте.
		if val > 1 {
			res = append(res, key) //пишем пересечения в слайс.
		}
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("пересечений нет")
	}

	return res, nil
}
