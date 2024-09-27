package main

import (
	"errors"
	"fmt"
	"strconv"
)

func arabToRome(num int) (string, error) {
	if num < 1 {
		return "", errors.New("Выдача паники, римские цифры не могут быть меньше единицы")
	}

	romeMap := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	roman := ""
	for value, symbol := range romeMap {
		for num >= value {
			num -= value
			roman += symbol
		}
	}
	return roman, nil
}

func romeToArab(roman string) (int, bool) {
	romeMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	val, exists := romeMap[roman]
	return val, exists
}

func toFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func checkRange(numA, numB float64) bool {
	return numA >= 1 && numA <= 10 && numB >= 1 && numB <= 10
}

func calculate(numA, numB float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return numA + numB, nil
	case "-":
		return numA - numB, nil
	case "*":
		return numA * numB, nil
	case "/":
		return numA / numB, nil
	default:
		return 0, errors.New("Выдача паники, неизвестная операция.")
	}
}

func main() {
	var a, b, operation string
	fmt.Scanln(&a, &operation, &b)

	// Флаги для отслеживания систем счисления
	isRomanA, isRomanB := false, false

	// Преобразование первого операнда
	if valA, ok := romeToArab(a); ok {
		a = fmt.Sprint(valA)
		isRomanA = true
	}

	// Преобразование второго операнда
	if valB, ok := romeToArab(b); ok {
		b = fmt.Sprint(valB)
		isRomanB = true
	}

	// Проверка на использование разных систем счисления
	if isRomanA != isRomanB {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
		return
	}

	// Преобразование строк в числа
	numA, errA := toFloat(a)
	numB, errB := toFloat(b)

	if errA != nil || errB != nil {
		fmt.Println("Выдача паники, некорректный ввод.")
		return
	}

	// Проверка диапазона чисел
	if !checkRange(numA, numB) {
		fmt.Println("Выдача паники, числа должны быть в диапазоне от 1 до 10.")
		return
	}

	// Вычисление результата
	result, err := calculate(numA, numB, operation)
	if err != nil {
		fmt.Println("Выдача паники, ", err, ".")
		return
	}

	// Вывод результата
	if isRomanA && isRomanB {
		if intResult, err := arabToRome(int(result)); err == nil {
			fmt.Println(intResult)
		} else {
			fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
	} else {
		fmt.Println(result)
	}
}
