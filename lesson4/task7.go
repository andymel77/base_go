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
7. Создать (не программно) текстовый файл, в котором каждая строка должна содержать данные о фирме: название, форма собственности, выручка, издержки.
Пример строки файла: firm_1 ООО 10000 5000.
Необходимо построчно прочитать файл, вычислить прибыль каждой компании, а также среднюю прибыль. Если фирма получила убытки, в расчет средней прибыли ее не включать.
Далее реализовать список. Он должен содержать словарь с фирмами и их прибылями, а также словарь со средней прибылью. Если фирма получила убытки, также добавить ее в словарь (со значением убытков).
Пример списка: [{“firm_1”: 5000, “firm_2”: 3000, “firm_3”: 1000}, {“average_profit”: 2000}].
Пример json-объекта:
[{"firm_1": 5000, "firm_2": 3000, "firm_3": 1000}, {"average_profit": 2000}]
*/

func main() {
	_firms := make(map[string]float64)
	_avProfit := make(map[string]float64)

	_json := []map[string]float64{}

	file, err := os.Open("lesson4/task7.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numFirms := 0
	sum := 0.0

	for scanner.Scan() {
		numFirms++

		_firmData := strings.Split(scanner.Text(), " ")
		_income, err := strconv.ParseFloat(_firmData[2], 10)

		if err != nil {
			log.Fatal(err)
		}

		_consumption, err := strconv.ParseFloat(_firmData[3], 10)

		if err != nil {
			log.Fatal(err)
		}

		profit := _income - _consumption
		sum += profit

		_firms[_firmData[0]] = profit
	}

	_avProfit["average_profit"] = sum / float64(numFirms)

	_json = append(_json, _firms)
	_json = append(_json, _avProfit)

	fmt.Println(_json)
}
