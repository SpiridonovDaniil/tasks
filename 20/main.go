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
	str := strings.Fields(text)      //делим строку по пробелу и пишем их в слайс строк.

	res := reverse(str) //делаем реверс слайса строк.

	var resString string //инициализируем результирующую строку.
	for _, val := range res {
		resString += val + " " //пишем в строку данные из слайса через пробел.
	}

	resString = strings.Trim(resString, " ") //обрезаем лишние пробелы по краям строки.
	fmt.Println(resString)                   //пишем в стандартный поток вывода получившуюся строку.
}

func reverse(arr []string) []string {
	j := len(arr) - 1
	for i := 0; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
	return arr
}
