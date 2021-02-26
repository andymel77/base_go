package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
1. Реализовать скрипт, в котором должна быть предусмотрена функция расчета заработной платы сотрудника.
В расчете необходимо использовать формулу: (выработка в часах * ставка в час) + премия. Для выполнения
расчета для конкретных значений необходимо запускать скрипт с параметрами.
*/
func main() {
	workHours, err := strconv.ParseFloat(os.Args[1], 10)
	if err != nil {
		log.Fatal(err)
	}
	ratePerHour, err := strconv.ParseFloat(os.Args[2], 10)
	if err != nil {
		log.Fatal(err)
	}
	prize, err := strconv.ParseFloat(os.Args[3], 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Salary:", workHours*ratePerHour+prize)
}
