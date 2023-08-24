package main

import (
	"fmt"
	"log"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Println(err)
	}
	if str == "" {
		log.Println("введите корректную строку")
		return
	}
	var x, y, z, cl int
	r := []rune(str)
	slo := make([]rune, 0)

	for _, val := range r {
		if x == -1 || y == -1 || z == -1 {
			break
		}

		switch {
		case val == '[':
			x++
			slo = append(slo, '[')
			cl--
		case val == ']':
			x--
			cl++
			if cl > 0 {
				fmt.Println("Скобки несбалансированы")
				return
			} else if slo[len(slo)-1] != '[' {
				fmt.Println("Скобки несбалансированы")
				return
			} else {
				slo = slo[:len(slo)-1]
			}
		case val == '(':
			y++
			slo = append(slo, '(')
			cl--
		case val == ')':
			y--
			cl++
			if cl > 0 {
				fmt.Println("Скобки несбалансированы")
				return
			} else if slo[len(slo)-1] != '(' {
				fmt.Println("Скобки несбалансированы")
				return
			} else {
				slo = slo[:len(slo)-1]
			}
		case val == '{':
			z++
			slo = append(slo, '{')
			cl--
		case val == '}':
			z--
			cl++
			if cl > 0 {
				fmt.Println("Скобки несбалансированы")
				return
			} else if slo[len(slo)-1] != '{' {
				fmt.Println("Скобки несбалансированы")
				return
			} else {
				slo = slo[:len(slo)-1]
			}
		}
	}

	if x == 0 && y == 0 && z == 0 {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}
}
