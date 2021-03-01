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
3. Создать текстовый файл (не программно), построчно записать фамилии сотрудников и величину их окладов.
Определить, кто из сотрудников имеет оклад менее 20 тыс., вывести фамилии этих сотрудников.
Выполнить подсчет средней величины дохода сотрудников.
*/

func main() {
	file, err := os.Open("lesson4/task3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numStr := 0
	sumSalary := 0.0

	for scanner.Scan() {
		numStr++

		textString := scanner.Text()
		salary, err := strconv.ParseFloat(strings.Split(textString, " ")[1], 10)

		if err != nil {
			log.Fatal(err)
		}

		employee := strings.Split(textString, " ")[0]

		if salary < 20000 {
			fmt.Println("Employee", employee, "has salary less 20000")
		}
		sumSalary += salary
	}
	fmt.Println("Average salary is: ", sumSalary/float64(numStr))
}
