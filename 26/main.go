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

	fmt.Println(check1(text)) //вариант с картой. Самый быстрый метод, так как осуществляется лишь один проход по строке.
	fmt.Println(check2(text)) //вариант с strings.ContainsRune.
	fmt.Println(check3(text)) //вариант с strings.Contains и strings.Cut, отличающийся от предыдущего лишь тем, что используем встроенный метод для вырезания элемента.
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
		sl := append([]rune(text)[:idx], []rune(text)[idx+1:]...) //копируем слайс рун строки без искомого элемента.
		if strings.ContainsRune(string(sl), val) {                //ищем в слайсе еще один искомый элемент.
			return false
		}
	}

	return true
}

func check3(text string) bool {
	for _, val := range []rune(text) { //итерируемся по строке в виде рун(для корректной работы с символами unicode).
		before, after, _ := strings.Cut(text, string(val)) //получаем строки до элемента val и после.
		if strings.Contains(before+after, string(val)) {   //ищем элемент в строке before+after.
			return false
		}
	}

	return true
}

//также как и в варинте 2 и 3 можно после вырезки использовать strings.IndexAny() или strings.IndexRune, но по своей сути логика такая же.
