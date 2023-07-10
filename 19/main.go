package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin) //создаем сканер для чтения из стандартного поток ввода.
	sc.Scan()                        //читаем из стандартного ввода в сканнер.
	text := sc.Text()                //считываем переводим данные из байта в строку.

	res := reverse(text)     //делаем реверс слайса рун.
	fmt.Println(string(res)) //пишем результат в стандартный поток вывода.
}

func reverse(text string) []rune {
	arr := []rune(text)
	j := len(arr) - 1
	for i := 0; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
	return arr
}
