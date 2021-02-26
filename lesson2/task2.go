package main

import "fmt"

/*
2. Реализовать функцию, принимающую несколько параметров, описывающих данные пользователя: имя, фамилия, год рождения, город проживания, email, телефон. Функция должна принимать параметры как именованные аргументы. Реализовать вывод данных о пользователе одной строкой.
*/

func userData(firstName, lastName, birthDate, city, email, phone string) {
	fmt.Printf("First name: %s; Last name: %s; Birth date: %s; City: %s; E-mail: %s; Phone number: %s\n", firstName, lastName, birthDate, city, email, phone)
}

func main() {
	userData("Ivan", "Ivanov", "20/01/1981", "Moscow", "ivanov.i@gmail.com", "+79889889988")
}
