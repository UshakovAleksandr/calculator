package main

import (
	"bufio"
	"fmt"
	"os"

	"calculator/funcs"
)

func main() {
	const add string = "+"
	const sub string = "-"
	const mult string = "*"
	const div string = "/"

	reader := bufio.NewReader(os.Stdin)

	var firstNumber float64
	var secondNumber float64
	var mathSign string

	for {
		number, err := funcs.ConvertToInt(funcs.InputNumber(*reader))
		if err != nil {
			fmt.Println("Введено неверное значение, повторите ввод")
			continue
		}
		firstNumber = number
		break
	}

	for {
		mathSign = funcs.InputMathSign(*reader)
		if funcs.CheckMathSign(mathSign, add, sub, mult, div) {
			break
		}
		fmt.Println("Введено неверное значение, повторите ввод")
		continue
	}

	for {
		number, err := funcs.ConvertToInt(funcs.InputNumber(*reader))
		if err != nil {
			fmt.Println("Введено неверное значение, повторите ввод")
			continue
		}
		secondNumber = number
		break
	}

	switch mathSign {
	case add:
		fmt.Printf("Результат сложения: %v\n", firstNumber+secondNumber)
	case sub:
		fmt.Printf("Результат вычитания: %v\n", firstNumber-secondNumber)
	case mult:
		fmt.Printf("Результат умножения: %v\n", firstNumber*secondNumber)
	case div:
		if secondNumber == 0 {
			fmt.Println("Деление на 0 запрещено")
		} else {
			fmt.Printf("Результат деления: %v\n", firstNumber/secondNumber)
		}
	}
}
