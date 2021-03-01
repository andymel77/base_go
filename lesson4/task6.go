package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
6. Необходимо создать (не программно) текстовый файл, где каждая строка описывает учебный предмет и наличие лекционных, практических и лабораторных занятий по этому предмету и их количество. Важно, чтобы для каждого предмета не обязательно были все типы занятий. Сформировать словарь, содержащий название предмета и общее количество занятий по нему. Вывести словарь на экран.
Примеры строк файла:
Информатика: 100(л) 50(пр) 20(лаб).
Физика: 30(л) — 10(лаб)
Физкультура: — 30(пр) —
Пример словаря:
{“Информатика”: 170, “Физика”: 40, “Физкультура”: 30}
*/

func main() {

	subjectMap := make(map[string]float64)

	file, err := os.Open("lesson4/task6.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d+[0-9]`)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		strText := scanner.Text()
		subject := strings.Split(strText, ":")[0]
		_marks := re.FindAll([]byte(strText), -1)

		sum := 0.0
		for _, el := range _marks {
			_el, err := strconv.ParseFloat(string(el), 10)

			if err != nil {
				log.Fatal(err)
			}
			sum += _el
		}

		subjectMap[subject] = sum
	}

	fmt.Println(subjectMap)
}
