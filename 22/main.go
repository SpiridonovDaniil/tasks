package main

import (
	"fmt"
	"math/big"
)

func main() {
	//используем пакет Big.
	a := big.NewInt(435234693523)
	b := big.NewInt(3429523745)
	z := big.NewInt(0)
	z.Mul(a, b)
	fmt.Println("умножение", z)
	z.Div(a, b)
	fmt.Println("деление", z)
	z.Add(a, b)
	fmt.Println("сложение", z)
	z.Sub(a, b)
	fmt.Println("вычитание", z)
}
