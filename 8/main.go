package main

import (
	"fmt"
	"log"
)

func main() {
	var n, i, b int64
	_, err := fmt.Scan(&n, &i, &b)
	if err != nil {
		log.Fatalln(err)
	}

	changeBit(i, b, n)
}

func changeBit(i, b, num int64) int {
	return num & (b << i)
}
