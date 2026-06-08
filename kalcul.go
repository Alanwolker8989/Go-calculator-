package main

import (
	"fmt"
)

func main() {
	var num1, num2 float32
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите оператор: ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Println("Результат:", num1+num2)
	case "-":
		fmt.Println("Результат: ", num1-num2)
	case "*":
		fmt.Println("Результат: ", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя!!!")
		} else {
			fmt.Println("Результат: ", num1/num2)
		}
	case "%":
		if int(num2) == 0 {
			fmt.Println("Делить на ноль нельзя!!!")
		} else {
			result := int(num1) % int(num2)
			fmt.Println("Результат: ", result)
		}
	default:
		fmt.Println("Неизвестная операция(")

	}
}
