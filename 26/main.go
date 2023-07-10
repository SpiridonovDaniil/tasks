package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin) //создаем сканер для чтения из стандартного поток ввода.
	sc.Scan()                        //читаем из стандартного ввода в сканнер.
	text := sc.Text()                //считываем переводим данные из байта в строку.
	text = strings.ToLower(text)     //переводим все элементы строки в нижний регистр, так как функция проверки должна быть регистронезависимой.

	fmt.Println(check1(text)) //вариант с картой. Самый быстрый метод, так как осуществляется лишь один проход по строке, но выделяет много памяти.
	fmt.Println(check2(text)) //вариант с strings.ContainsRune.
}

func check1(text string) bool {
	data := make(map[rune]int)         //инициализируем карту, в которой будем считать сколько значий равных ключу встретится в слайсе.
	for _, val := range []rune(text) { //итерируемся по строке в виде рун(для корректной работы с символами unicode).
		data[val] += 1
		if data[val] > 1 { //если значение в карте превысило 1, значит симол не уникальный.
			return false
		}
	}

	return true
}

func check2(text string) bool {
	for idx, val := range []rune(text) { //итерируемся по строке в виде рун(для корректной работы с символами unicode).
		if strings.ContainsRune(string([]rune(text)[idx+1:]), val) { //ищем в последующей части слайса еще один искомый элемент(val).
			return false
		}
	}

	return true
}
