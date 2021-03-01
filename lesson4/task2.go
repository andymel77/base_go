package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
2. Создать текстовый файл (не программно), сохранить в нем несколько строк, выполнить подсчет количества строк, количества слов в каждой строке.
*/
func main() {
	file, err := os.Open("lesson4/task2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numStr := 0

	for scanner.Scan() {
		numStr++
		fmt.Println("'", scanner.Text(), "' has", len(strings.Split(scanner.Text(), " ")), "word(s)")
	}
	fmt.Println("Number of strings are: ", numStr)
}
