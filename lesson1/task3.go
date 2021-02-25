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
3. Пользователь вводит месяц в виде целого числа от 1 до 12. Сообщить к какому времени года относится месяц (зима, весна, лето, осень).
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

func inSeasons(season []int, month int) bool {

	for _, s := range season {
		if month == s {
			return true
		}
	}

	return false
}

func main() {
	winter := []int{12, 1, 2}
	spring := []int{3, 4, 5}
	summer := []int{6, 7, 8}
	autumn := []int{9, 10, 11}

	month, err := readIntFromInput("Input month: ")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if month < 1 || month > 12 {
		log.Fatal("Month not in range from 1 to 12")
		os.Exit(1)
	}

	if inSeasons(winter, month) {
		fmt.Println("It's winter now")
	} else if inSeasons(spring, month) {
		fmt.Println("It's spring now")
	} else if inSeasons(summer, month) {
		fmt.Println("It's summer now")
	} else if inSeasons(autumn, month) {
		fmt.Println("It's autumn now")
	}
}
