package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print("Введите выражение: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		expression := input.Text()

		firstNum, secondNum, isRome := NumCheck(expression)

		str := strings.Split(expression, " ")
		operator := str[1]

		result := calc(firstNum, secondNum, operator)
		if isRome && result > 0 {
			resultRome, _ := ArabToRome(result)
			fmt.Println(resultRome)
		} else if isRome && result <= 0 {
			panic("В римской системе счисления результат должен быть больше 0")
		} else {
			fmt.Println("Ответ:", result)
		}
	}
}

func NumCheck(expression string) (int, int, bool) {
	str := strings.Split(expression, " ")

	if len(str) != 3 {
		panic("Неверный формат математической операции")
	}

	isRome1 := RomeOrArab(str[0])
	isRome2 := RomeOrArab(str[2])

	var firstNum, secondNum int
	var err error

	if isRome1 && isRome2 {
		firstNum, err = RomeToArab(str[0])
		if err != nil {
			panic("Неверный формат чисел")
		}
		secondNum, err = RomeToArab(str[2])
		if err != nil {
			panic("Неверный формат чисел")
		}
	} else if !isRome1 && !isRome2 {
		firstNum, err = strconv.Atoi(str[0])
		if err != nil {
			panic("Неверный формат чисел")
		}
		secondNum, err = strconv.Atoi(str[2])
		if err != nil {
			panic("Неверный формат чисел")
		}
	} else {
		panic("Допускается ввод только арабских или только римских чисел")
	}

	if firstNum > 10 || firstNum < 1 || secondNum > 10 || secondNum < 1 {
		panic("Числа должны быть от 1 до 10")
	}
	return firstNum, secondNum, isRome1
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

func ArabToRome(num int) (string, error) {
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

	keys := make([]int, 0, len(romeMap))
	for k := range romeMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var result strings.Builder
	for _, k := range keys {
		for num >= k {
			result.WriteString(romeMap[k])
			num -= k
		}
	}

	return result.String(), nil
}

func RomeOrArab(s string) bool {
	if s == "I" || s == "II" || s == "III" || s == "IV" || s == "V" || s == "VI" || s == "VII" || s == "VIII" || s == "IX" || s == "X" {
		return true
	}
	return false
}
