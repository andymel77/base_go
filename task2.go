package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
2. Для списка реализовать обмен значений соседних элементов, т.е. Значениями обмениваются элементы с индексами 0 и 1, 2 и 3 и т.д.
При нечетном количестве элементов последний сохранить на своем месте. Для заполнения списка элементов необходимо использовать функцию input().
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
	n, err := readIntFromInput("Enter number of elements: ")

	if err != nil {
		log.Fatal(err)
		return
	}

	numbers := []int{}

	for i := 0; i < n; i++ {
		num, err := readIntFromInput("Enter number #" + fmt.Sprint(i+1) + ": ")

		if err != nil {
			log.Fatal(err)
			return
		}

		numbers = append(numbers, num)
	}

	fmt.Println("Original:", numbers)

	for i := 0; i < n; i++ {
		if (i+1)%2 == 0 {
			numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
		}
	}

	fmt.Println("Changed: ", numbers)
}
