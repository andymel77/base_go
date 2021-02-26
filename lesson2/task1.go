package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
1. Реализовать функцию, принимающую два числа (позиционные аргументы) и выполняющую их деление. Числа запрашивать у пользователя, предусмотреть обработку ситуации деления на ноль.
*/

func readFloatFromInput(message string) (float64, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	ns, _ := reader.ReadString('\n')
	ns = strings.Replace(ns, "\n", "", -1)

	n, err := strconv.ParseFloat(ns, 10)

	if err != nil {
		return 0, err
	}

	return float64(n), nil
}

func division(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}

func main() {
	a, err := readFloatFromInput("Enter number a: ")

	if err != nil {
		log.Fatal(err)
	}

	b, err := readFloatFromInput("Enter number b: ")

	if err != nil {
		log.Fatal(err)
	}

	res, err := division(a, b)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
