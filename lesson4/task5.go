package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

/*
5. Создать (программно) текстовый файл, записать в него программно набор чисел, разделенных пробелами.
Программа должна подсчитывать сумму чисел в файле и выводить ее на экран.
*/
func main() {
	file, err := os.OpenFile("lesson4/task5.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	source := rand.NewSource(time.Now().UnixNano())
	_rand := rand.New(source)
	_res := []int{}

	for i := 0; i < 20; i++ {
		randNum := _rand.Intn(1000)
		_res = append(_res, randNum)
		file.WriteString(fmt.Sprint(randNum) + " ")
	}

	file.WriteString("\n")

	sum := 0

	for _, el := range _res {
		sum += el
	}

	fmt.Println("Sum of numbers: ", sum)
}
