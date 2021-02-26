package main

import (
	"fmt"
	"sort"
)

/*
3. Реализовать функцию my_func(), которая принимает три позиционных аргумента, и возвращает сумму наибольших двух аргументов.
*/

func myFunc(a, b, c int) int {
	_slice := []int{a, b, c}

	sort.Slice(_slice, func(i, j int) bool {
		return _slice[i] > _slice[j]
	})

	return _slice[0] + _slice[1]
}

func main() {
	fmt.Println(myFunc(4, 3, 5))
}
