package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
4. Создать (не программно) текстовый файл со следующим содержимым:
One — 1
Two — 2
Three — 3
Four — 4
Необходимо написать программу, открывающую файл на чтение и считывающую построчно данные.
При этом английские числительные должны заменяться на русские. Новый блок строк должен записываться в новый текстовый файл.
*/

func main() {
	numberMap := make(map[string]string)

	numberMap["1"] = "Один"
	numberMap["2"] = "Два"
	numberMap["3"] = "Три"
	numberMap["4"] = "Четыре"

	_result := []string{}

	file, err := os.Open("lesson4/task4_in.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strText := scanner.Text()

		numbers := strings.Split(strText, " ")
		_resString := numbers[0] + " - " + numberMap[numbers[2]]
		_result = append(_result, _resString)
	}

	fileOut, err := os.OpenFile("lesson4/task4_out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer fileOut.Close()

	for _, el := range _result {
		fileOut.WriteString(el + "\n")
	}
}
