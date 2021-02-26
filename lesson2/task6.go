package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
6. Реализовать функцию int_func(), принимающую слово из маленьких латинских букв и возвращающую его же, но с прописной первой буквой. Например, print(int_func(‘text’)) -> Text.
Продолжить работу над заданием. В программу должна попадать строка из слов, разделенных пробелом. Каждое слово состоит из латинских букв в нижнем регистре.
Сделать вывод исходной строки, но каждое слово должно начинаться с заглавной буквы. Необходимо использовать написанную ранее функцию int_func().
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

func intFunc(text string) string {
	return strings.Title(text)
}

func main() {
	str, err := readStringFromInput("Enter text: ")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(intFunc(str))
}
