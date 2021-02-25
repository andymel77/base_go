package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
5. Реализовать структуру «Рейтинг», представляющую собой не возрастающий набор натуральных чисел. У пользователя необходимо запрашивать новый элемент рейтинга. Если в рейтинге существуют элементы с одинаковыми значениями, то новый элемент с тем же значением должен разместиться после них.
Подсказка. Например, набор натуральных чисел: 7, 5, 3, 3, 2.
Пользователь ввел число 3. Результат: 7, 5, 3, 3, 3, 2.
Пользователь ввел число 8. Результат: 8, 7, 5, 3, 3, 2.
Пользователь ввел число 1. Результат: 7, 5, 3, 3, 2, 1.
Набор натуральных чисел можно задать непосредственно в коде, например, my_list = [7, 5, 3, 3, 2].
*/

func readIntFromInput(message string) (int, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	ns, _ := reader.ReadString('\n')
	ns = strings.Replace(ns, "\n", "", -1)

	n, err := strconv.ParseInt(ns, 10, 32)

	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func main() {
	myList := []int{7, 5, 3, 3, 2}

	n, err := readIntFromInput("Enter a number: ")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	myList = append(myList, n)

	fmt.Println("Origin list:", myList)

	sort.Slice(myList, func(i, j int) bool {
		return myList[i] > myList[j]
	})

	fmt.Println("Sorted list:", myList)
}
