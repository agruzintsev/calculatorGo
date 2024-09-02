package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("Введите выражение: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		expression := input.Text()
		firstNum, secondNum := NumCheck(expression)

		str := strings.Split(expression, " ")
		operator := str[1]

		result := calc(firstNum, secondNum, operator)
		fmt.Println(result)
	}
}

func NumCheck(expression string) (int, int) {
	str := strings.Split(expression, " ")

	isRome1, _ := RomeOrArab(str[0])
	isRome2, _ := RomeOrArab(str[2])

	var firstNum, secondNum int
	var err error

	if isRome1 && isRome2 {
		firstNum, err = RomeToArab(str[0])
		if err != nil {
			panic(err)
		}
		secondNum, err = RomeToArab(str[2])
		if err != nil {
			panic(err)
		}
	} else if !isRome1 && !isRome2 {
		firstNum, err = strconv.Atoi(str[0])
		if err != nil {
			panic(err)
		}
		secondNum, err = strconv.Atoi(str[2])
		if err != nil {
			panic(err)
		}
	} else {
		panic("Только арабские или только римские цифры")
	}

	if firstNum > 10 || firstNum < 1 || secondNum > 10 || secondNum < 1 {
		panic("Числа должны быть от 1 до 10")
	}
	return firstNum, secondNum
}

func calc(firstNum, secondNum int, operator string) int {
	switch operator {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	default:
		panic("Неизвестный оператор")
	}
}

func RomeToArab(rome string) (int, error) {
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
	return romeMap[rome], nil
}

func RomeOrArab(s string) (bool, error) {
	if s == "I" || s == "II" || s == "III" || s == "IV" || s == "V" || s == "VI" || s == "VII" || s == "VIII" || s == "IX" || s == "X" {
		return true, nil
	}
	return false, nil
}
