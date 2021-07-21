package funcs

import (
	"fmt"
	"strconv"
)

func ConvertToInt(num string) (float64, error) {
	number, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Printf("Введенный символ '%v' не является числом. Ошибка: '%v'\n", num, err)
	}

	return number, err
}

func CheckMathSign(mathSign string, add string, sub string, mult string, div string) bool {
	avalibleSigns := [4]string{add, sub, mult, div}

	var check bool = false

	for _, sign := range avalibleSigns {
		if mathSign == sign {
			check = true
			break
		}
	}
	return check
}
