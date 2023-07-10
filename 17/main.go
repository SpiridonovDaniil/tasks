package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	m := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	r := sort.SearchInts(m, 6)
	fmt.Println("Элемент 6 не найден, но мог бы быть вставлен по индексу:", r)
	r = sort.SearchInts(m, 5)
	fmt.Println("Элемент 5 расположен по индексу:", r)

	s := []string{"a", "b", "c", "e", "f"}
	rs := sort.SearchStrings(s, "d")
	fmt.Println("Элемент d не найден, но мог бы быть вставлен по индексу:", rs)
	rs = sort.SearchStrings(s, "f")
	fmt.Println("Элемент f расположен по индексу:", rs)

	f := []float64{1.1, 1.2, 1.3, 1.4, 1.7, 1.8, 1.9}
	rf := sort.SearchFloat64s(f, 1.5)
	fmt.Println("Элемент 1.5 не найден, но мог бы быть вставлен по индексу:", rf)
	rf = sort.SearchFloat64s(f, 1.2)
	fmt.Println("Элемент 1.2 расположен по индексу:", rf)

	a := []int{1, 3, 6, 6, 7, 8, 9}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x }) //указываем, что поиск ведется в отсортированном по возрастанию слайсе
	if i < len(a) && a[i] == x {
		fmt.Println("элемент 6 найден по индексу", i) //стоит обратить внимание, что поиск возвращает первое вхождение объекта в слайсе
	} else {
		fmt.Println("элемент 6 не найден, но мог бы быть вставлен по индексу:", i)
	}

	y := []string{"a", "b", "c", "d"}
	target := "a"
	z, found := sort.Find(y.Len(), func(i int) int {
		return strings.Compare(target, y.At(z))
	})
	if found {
		fmt.Printf("found %s at entry %d\n", target, z)
	} else {
		fmt.Printf("%s not found, would insert at %d", target, z)
	}
	//TODO попробывать реализовать функцию Find
}
