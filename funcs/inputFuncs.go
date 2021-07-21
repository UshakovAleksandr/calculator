package funcs

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func InputNumber(reader bufio.Reader) string {
	fmt.Print("Введите число: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)

	return input
}

func InputMathSign(reader bufio.Reader) string {
	fmt.Print("Введите мат. знак: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)

	return input
}
