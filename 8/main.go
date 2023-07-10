package main

import (
	"fmt"
	"log"
)

var n int64

func main() {
	var i, b int
	_, err := fmt.Scan(&i, &b)
	if err != nil {
		log.Fatalln(err)
	}

	changeBit(i, b, n)
}

func changeBit(i, b int, num int64) {
	укпф
}
