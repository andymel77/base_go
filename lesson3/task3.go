package main

import "fmt"

/*
3. Для чисел в пределах от 20 до 240 найти числа, кратные 20 или 21. Необходимо решить задание в одну строку.
*/

func generateIntSlice(start, end int) []int {
	_nums := []int{}

	for i := start; i <= end; i++ {
		_nums = append(_nums, i)
	}

	return _nums
}

func main() {
	nums := generateIntSlice(20, 240)

	_res := []int{}

	for _, el := range nums {
		if el%20 == 0 || el%21 == 0 {
			_res = append(_res, el)
		}
	}

	fmt.Println(_res)
}
