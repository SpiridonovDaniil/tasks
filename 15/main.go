package main

import (
	"fmt"
	"math/rand"
)

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10) //512 знаков??
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
	v := createHugeString(1 << 10)
	v = "ущпрцзтщушкепртукщпртк"
	justString = string([]rune(v)[:100])
	fmt.Println(justString)
}

func createHugeString(num int) string {
	arr := make([]rune, num)
	for i := 0; i < num; i++ {
		// todo pakcage unicode https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
		arr[i] = rand.Int31n(440)
	}

	return string(arr)
}

//данная операция v[:100] может привести к проблеме если символы строки кодируются больше чем одним байтом.
//к примеру если первые 50 символов будут закодированы 2 байтами, то на выходе мы получим 50 символов строки.
