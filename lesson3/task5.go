package main

import "fmt"

/*
5. Реализовать формирование списка, используя функцию range() и возможности генератора.
В список должны войти четные числа от 100 до 1000 (включая границы). Необходимо получить результат вычисления произведения всех элементов списка.
*/

func generateIntSlice(start, end int) []int {
	_nums := []int{}

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			_nums = append(_nums, i)
		}
	}

	return _nums
}

func main() {
	_nums := generateIntSlice(100, 110)

	var res int64 = 1

	for _, el := range _nums {
		res *= int64(el)
	}

	fmt.Println("Mult:", res)
}
