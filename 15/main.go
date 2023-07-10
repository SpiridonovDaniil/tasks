package main

import (
	"fmt"
)

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//func main() {
//	someFunc()
//}

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10) //создаем строку из 1024 символов.

	justString = string([]rune(v)[:100]) //берем из созданной строки первые 100 рун.
	fmt.Println(justString)
}

func createHugeString(num int) string {
	arr := make([]rune, num)
	for i := 0; i < num; i++ {
		arr[i] = int32(i + 1040) //намерено заполняем строку рунами, закодированными 2 байтами(256-65000).
	}

	return string(arr)
}

//данная операция v[:100] может привести к проблеме если символы строки кодируются больше чем одним байтом.
//к примеру если первые 50 символов будут закодированы 2 байтами, то на выходе мы получим 50 символов строки.
