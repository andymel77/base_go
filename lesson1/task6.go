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
6. * Реализовать структуру данных «Товары». Структуру нужно сформировать программно, т.е. запрашивать все данные у пользователя.
Пример готовой структуры:
[
(1, {“название”: “компьютер”, “цена”: 20000, “количество”: 5, “eд”: “шт.”}),
(2, {“название”: “принтер”, “цена”: 6000, “количество”: 2, “eд”: “шт.”}),
(3, {“название”: “сканер”, “цена”: 2000, “количество”: 7, “eд”: “шт.”})
]
Необходимо собрать аналитику о товарах. Реализовать словарь, в котором каждый ключ — характеристика товара, например название, а значение — список значений-характеристик, например список названий товаров.
Пример:
{
“название”: [“компьютер”, “принтер”, “сканер”],
“цена”: [20000, 6000, 2000],
“количество”: [5, 2, 7],
“ед”: [“шт.”]
}
*/

type hard struct {
	name  string
	price float64
	count float64
	unit  string
}

type hardChars struct {
	names  []string
	prices []float64
	counts []float64
	units  []string
}

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

func (h hardChars) findAtNames(name string) bool {

	for _, n := range h.names {
		if name == n {
			return true
		}
	}

	return false
}

func (h hardChars) findAtPrices(price float64) bool {

	for _, p := range h.prices {
		if price == p {
			return true
		}
	}

	return false
}

func (h hardChars) findAtCounts(count float64) bool {

	for _, c := range h.counts {
		if count == c {
			return true
		}
	}

	return false
}

func (h hardChars) findAtUnits(unit string) bool {

	for _, u := range h.units {
		if unit == u {
			return true
		}
	}

	return false
}

func main() {
	hards := []hard{}

	i := 1
	for {
		name, err := readStringFromInput("Enter a hard " + fmt.Sprint(i) + " name: ")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		price, err := readFloatFromInput("Enter a hard " + fmt.Sprint(i) + " price: ")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		count, err := readFloatFromInput("Enter a hard " + fmt.Sprint(i) + " count: ")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		unit, err := readStringFromInput("Enter a hard " + fmt.Sprint(i) + " unit: ")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		_hard := hard{name, price, count, unit}

		hards = append(hards, _hard)

		answer, err := readStringFromInput("Enter one more new hard (YES, no) ")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if "no" == strings.ToLower(answer) {
			break
		}

		i++
	}

	fmt.Println(hards)

	hardChar := hardChars{}

	for _, h := range hards {
		if !hardChar.findAtNames(h.name) {
			hardChar.names = append(hardChar.names, h.name)
		}

		if !hardChar.findAtPrices(h.price) {
			hardChar.prices = append(hardChar.prices, h.price)
		}

		if !hardChar.findAtCounts(h.count) {
			hardChar.counts = append(hardChar.counts, h.count)
		}

		if !hardChar.findAtUnits(h.unit) {
			hardChar.units = append(hardChar.units, h.unit)
		}
	}

	fmt.Println(hardChar)
}
