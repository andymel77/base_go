package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
4. Пользователь вводит строку из нескольких слов, разделённых пробелами. Вывести каждое слово с новой строки. Строки необходимо пронумеровать. Если слово длинное, выводить только первые 5 букв в слове.
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
	str, err := readStringFromInput("Enter your string: ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)

	strSlice := strings.Split(str, " ")

	for i, ss := range strSlice {
		if len(ss) > 10 {
			fmt.Println("String #", i+1, ":", string([]rune(ss)[0:10])+"")
		} else {
			fmt.Println("String #", i+1, ":", ss)
		}
	}
}
