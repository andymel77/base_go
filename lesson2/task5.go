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
5. Программа запрашивает у пользователя строку чисел, разделенных пробелом. При нажатии Enter должна выводиться сумма чисел.
Пользователь может продолжить ввод чисел, разделенных пробелом и снова нажать Enter. Сумма вновь введенных чисел будет добавляться
к уже подсчитанной сумме. Но если вместо числа вводится специальный символ, выполнение программы завершается. Если специальный
символ введен после нескольких чисел, то вначале нужно добавить сумму этих чисел к полученной ранее сумме и после этого завершить программу.
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
	summa := 0.0
	exitSymbol := false

	for {
		ns, err := readStringFromInput("Enter numbers separated by space (special symbol # to stop): ")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		_slice := strings.Split(ns, " ")

		for _, el := range _slice {
			if el == "#" {
				exitSymbol = true
				break
			}

			_num, err := strconv.ParseFloat(el, 10)

			if err != nil {
				log.Fatal(err)
			}

			summa += _num
		}

		fmt.Println("Summa:", summa)

		if exitSymbol {
			break
		}
	}
}
