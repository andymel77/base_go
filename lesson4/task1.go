package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
1. Создать программно файл в текстовом формате, записать в него построчно данные, вводимые пользователем.
Об окончании ввода данных свидетельствует пустая строка.
*/

func readStringFromInput(message string) (string, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	ns, err := reader.ReadString('\n')
	ns = strings.Replace(ns, "\n", "", -1)

	if err != nil {
		return "", err
	}

	return ns, nil
}

func main() {
	file, err := os.OpenFile("task1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for {
		ns, err := readStringFromInput("Enter a string: ")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if ns == "" {
			break
		}

		file.WriteString(ns + "\n")
	}
}
