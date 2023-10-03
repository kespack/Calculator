package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertIntToRoman(number int) (string, error) {
	if number < 1 || number > 100 {
		return "", errors.New("Некорректная строка")
	}
	var num = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var sym = []string{"I", "IV", "V", "IX", "X", "XL",
		"L", "XC", "C"}
	i := len(num) - 1
	var roman strings.Builder
	for number != 0 {
		div := number / num[i]
		number %= num[i]

		for div != 0 {
			roman.WriteString(sym[i])
			div -= 1
		}
		i -= 1
	}
	return roman.String(), nil
}

func convertRomanToInt(roman string) (int, error) {
	var values = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100}
	roman = strings.ToUpper(roman)
	var res = 0

	for q := 0; q < len(roman); q++ {
		val1, ok1 := values[string(roman[q])]
		if !ok1 {
			return 0, errors.New("Некорректная строка")
		}
		if q == len(roman)-1 {
			res += val1
			break
		}
		val2, ok2 := values[string(roman[q+1])]
		if !ok2 {
			return 0, errors.New("Некорректная строка")
		}
		if val1 < val2 {
			res -= val1
		} else {
			res += val1
		}
	}

	return res, nil
}

func calculator(expression string) {
	expression = strings.ReplaceAll(expression, " ", "")
	if len(expression) < 3 {
		fmt.Println("Некорректная строка")
		return
	}
	var operations = []string{"+", "-", "/", "*"}
	var splitExpression []string
	var operation string

	for index := range operations {
		if strings.Contains(expression, operations[index]) {
			splitExpression = strings.Split(expression, operations[index])
			operation = operations[index]
		}
		for i := range splitExpression {
			splitExpression[i] = strings.TrimSpace(splitExpression[i])
		}
	}
	if len(splitExpression) != 2 {
		fmt.Println("Некорректная строка")
		return
	}
	var first int
	var second int
	var isRoman = false
	first, err := strconv.Atoi(splitExpression[0])
	if err == nil {
		second, err = strconv.Atoi(splitExpression[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		first, err = convertRomanToInt(splitExpression[0])
		if err == nil {
			second, err = convertRomanToInt(splitExpression[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			isRoman = true
		} else {
			fmt.Println(err)
			return
		}
	}

	if first < 1 || first > 10 || second < 1 || second > 10 {
		fmt.Println("Некорректная строка")
		return
	}

	var result = 0
	switch operation {
	case "+":
		result = first + second
		break
	case "-":
		result = first - second
		break
	case "/":
		result = first / second
		break
	case "*":
		result = first * second
		break
	}

	if isRoman {
		resultRoman, _ := convertIntToRoman(result)
		fmt.Println(resultRoman)
		return
	} else {
		fmt.Println(result)
		return
	}
}
func main() {
	for {
		in := bufio.NewReader(os.Stdin)
		expression, _ := in.ReadString('\n')
		if expression == "0" {
			return
		}
		calculator(expression)

	}
}
