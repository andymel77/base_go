package main

import "fmt"

/*
2. Реализовать функцию, принимающую несколько параметров, описывающих данные пользователя: имя, фамилия, год рождения, город проживания, email, телефон. Функция должна принимать параметры как именованные аргументы. Реализовать вывод данных о пользователе одной строкой.
*/

type Args struct {
	firstName, lastName, birthDate, city, email, phone string
}

func userData(firstName, lastName, birthDate, city, email, phone string) {
	fmt.Printf("First name: %s; Last name: %s; Birth date: %s; City: %s; E-mail: %s; Phone number: %s\n", firstName, lastName, birthDate, city, email, phone)
}

func userDataNamedParams(args *Args) {
	fmt.Printf("First name: %s; Last name: %s; Birth date: %s; City: %s; E-mail: %s; Phone number: %s\n", args.firstName, args.lastName, args.birthDate, args.city, args.email, args.phone)
}

func main() {
	userData("Ivan", "Ivanov", "20/01/1981", "Moscow", "ivanov.i@gmail.com", "+79889889988")
	userDataNamedParams(&Args{firstName: "Ivan", lastName: "Ivanov", birthDate: "20/01/1981", city: "Moscow", email: "ivanov.i@gmail.com", phone: "+79889889988"})
}
