package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// *******
// Функция определения наличия в заданном тексте "text" арифметических операторов: "+", "—", "*", "/".
// Если оператор в тексте "text" имеется, то функция возвращает оператор.
// Если оператора нет, или операторов больше одного, то будет выведена ошибка.
func operation(text string) string {
	var operator string
	var operators []string
	for i := range text {
		if rune(text[i]) == 43 {
			operators = append(operators, "+")
		}
		if rune(text[i]) == 45 {
			operators = append(operators, "-")
		}
		if rune(text[i]) == 42 {
			operators = append(operators, "*")
		}
		if rune(text[i]) == 47 {
			operators = append(operators, "/")
		}
	}
	if len(operators) == 1 {
		operator = operators[0]
		return operator
	}
	if len(operators) > 1 {
		return "Ошибка 1003, неправильный формат математической операции. Введите два арабских числа (от 1 до 10) или два римских числа (от I до X) и один оператор (+, -, /, *)."
	}
	return "Ошибка 1001, неправильный формат математической операции. Введите два арабских числа (от 1 до 10) или два римских числа (от I до X) и один оператор (+, -, /, *)."
}

// *******
// Функция разделения текста "text" по словам, которая возвращает массив [n], где n - количество элементов текста.
func textFields(text string) []string {
	arrWords := strings.Fields(text)
	if len(arrWords) <= 2 && len(arrWords) > 0 {
		return arrWords
	}
	return nil
}

// *******
// Функция, которая удаляет арифметический оператор в тексте "text", возвращая массив текста уже без оператора.
// Новая строка "textNew" передается в функцию "textFields", которая возвращает массив [n], где n - количество элементов текста.
func arrayLen(text string) []string {
	lookOperator := operation(text)
	lookSpace := " "
	arrText := strings.Split(text, lookOperator)
	arrTextSpace := strings.Split(text, lookSpace)
	switch lookOperator {
	case "+":
		if len(arrTextSpace) == 1 {
			return arrText
		} else {
			textNew := strings.Join(arrText, "") // Новый текст без оператора и пробелов.
			arrTextNew := textFields(textNew)    // Массив, состоящий из 2 элементов (числа).
			return arrTextNew
		}
	case "-":
		if len(arrTextSpace) == 1 {
			return arrText
		} else {
			textNew := strings.Join(arrText, "")
			arrTextNew := textFields(textNew)
			return arrTextNew
		}
	case "*":
		if len(arrTextSpace) == 1 {
			return arrText
		} else {
			textNew := strings.Join(arrText, "")
			arrTextNew := textFields(textNew)
			return arrTextNew
		}
	default:
		if len(arrTextSpace) == 1 {
			return arrText
		} else {
			textNew := strings.Join(arrText, "")
			arrTextNew := textFields(textNew)
			return arrTextNew
		}
	}
}

// *******
// Функция, которая переводит римские числа в арабские числа.
func romanToArabInt(text string) int {
	arr := text
	arr = arr + "   "
	var arrRomanToArab int

	for i := 0; i+1 < len(arr); {
		if rune(arr[i]) == 88 {
			if rune(arr[i+1]) == 67 {
				arrRomanToArab = arrRomanToArab + 90
				arr = arr[2:]
			}
			if rune(arr[i+1]) == 76 {
				arrRomanToArab = arrRomanToArab + 40
				arr = arr[2:]
			}
		}
		if rune(arr[i]) == 73 {
			if rune(arr[i+1]) == 88 {
				arrRomanToArab = arrRomanToArab + 9
				arr = arr[2:]
			}
			if rune(arr[i+1]) == 86 {
				arrRomanToArab = arrRomanToArab + 4
				arr = arr[2:]
			}
		}
		if rune(arr[i]) == 67 {
			arrRomanToArab = arrRomanToArab + 100
			arr = arr[1:]
		}
		if rune(arr[i]) == 76 {
			arrRomanToArab = arrRomanToArab + 50
			arr = arr[1:]

			if arr[i] == 76 {
				return 1006
			}
		}
		if rune(arr[i]) == 88 {
			arrRomanToArab = arrRomanToArab + 10
			arr = arr[1:]
			if rune(arr[i]) == 88 {
				arrRomanToArab = arrRomanToArab + 10
				arr = arr[1:]
				if rune(arr[i]) == 88 {
					arrRomanToArab = arrRomanToArab + 10
					arr = arr[1:]
					if rune(arr[i]) == 88 {
						return 1006
					}
				}
			}
		}
		if rune(arr[i]) == 86 {
			arrRomanToArab = arrRomanToArab + 5
			arr = arr[1:]
			if arr[i] == 86 {
				return 1006
			}
		}
		if rune(arr[i]) == 73 {
			arrRomanToArab = arrRomanToArab + 1
			arr = arr[1:]
			if rune(arr[i]) == 73 {
				arrRomanToArab = arrRomanToArab + 1
				arr = arr[1:]
				if rune(arr[i]) == 73 {
					arrRomanToArab = arrRomanToArab + 1
					arr = arr[1:]
					if rune(arr[i]) == 73 {
						return 1006
					}
				}
			}
		}
		return arrRomanToArab
	}
	return 1006
}

// *******
// Функция перевода римских чисел, которые содержатся в массиве из функции "arrayLen", в арабские числа.
func romanToArab(arrText []string) []string {
	var arrRomanToArab []string
	RomanToArab1 := romanToArabInt(arrText[0])
	RomanToArab2 := romanToArabInt(arrText[1])

	RomanToArab1String := strconv.Itoa(RomanToArab1)
	RomanToArab2String := strconv.Itoa(RomanToArab2)

	arrRomanToArab = append(arrRomanToArab, RomanToArab1String, RomanToArab2String)
	return arrRomanToArab
}

// *******
// Функция перевода арабских чисел в римские числа.
func arabToRoman(total int) string {
	totalText := strconv.Itoa(total)
	lookNil := ""
	totalNew := strings.Split(totalText, lookNil)

	if len(totalNew) == 3 {
		totalNew[0] = totalNew[0] + "00"
		totalNew[1] = totalNew[1] + "0"
		totalNew[2] = totalNew[2]
	}
	if len(totalNew) == 2 {
		totalNew[0] = totalNew[0] + "0"
		totalNew[1] = totalNew[1]
	}
	if len(totalNew) == 1 {
	}
	var roman string

	for i := range totalNew {
		if len(totalNew[i]) == 3 {
			if totalNew[i] == "100" {
				roman = roman + "C"
			}
		}
		if len(totalNew[i]) == 2 {
			if totalNew[i] == "10" {
				roman = roman + "X"
			}
			if totalNew[i] == "20" {
				roman = roman + "XX"
			}
			if totalNew[i] == "30" {
				roman = roman + "XXX"
			}
			if totalNew[i] == "40" {
				roman = roman + "XL"
			}
			if totalNew[i] == "50" {
				roman = roman + "L"
			}
			if totalNew[i] == "60" {
				roman = roman + "LX"
			}
			if totalNew[i] == "70" {
				roman = roman + "LXX"
			}
			if totalNew[i] == "80" {
				roman = roman + "LXXX"
			}
			if totalNew[i] == "90" {
				roman = roman + "XC"
			}
		}
		if len(totalNew[i]) == 1 {
			if totalNew[i] == "1" {
				roman = roman + "I"
			}
			if totalNew[i] == "2" {
				roman = roman + "II"
			}
			if totalNew[i] == "3" {
				roman = roman + "III"
			}
			if totalNew[i] == "4" {
				roman = roman + "IV"
			}
			if totalNew[i] == "5" {
				roman = roman + "V"
			}
			if totalNew[i] == "6" {
				roman = roman + "VI"
			}
			if totalNew[i] == "7" {
				roman = roman + "VII"
			}
			if totalNew[i] == "8" {
				roman = roman + "VIII"
			}
			if totalNew[i] == "9" {
				roman = roman + "IX"
			}
		}
	}
	return roman
}

// *******
// Функция, которая проводит математическую операцию с римскими числами.
func arrOperationRoman(text string, arr []string) string {
	arr = romanToArab(arr)

	x, err := strconv.Atoi(arr[0])
	if err != nil {
		return "1006"
	}
	y, err := strconv.Atoi(arr[1])
	if err != nil {
		return "1006"
	}
	if x > 0 && x <= 10 && y > 0 && y <= 10 {
		if operation(text) == "+" {
			total := x + y
			totalRoman := arabToRoman(total)
			return totalRoman
		}
		if operation(text) == "-" {
			total := x - y
			if total < 0 {
				return "1009"
			}
			totalRoman := arabToRoman(total)
			return totalRoman
		}
		if operation(text) == "/" {
			total := x / y
			totalRoman := arabToRoman(total)
			return totalRoman
		}
		if operation(text) == "*" {
			total := x * y
			totalRoman := arabToRoman(total)
			return totalRoman
		}
	}
	return "1006"
}

// *******
// Функция, которая проводит математическую операцию с арабскими числами.
func arrOperationArab(text string, arr []string) int {
	x, err := strconv.Atoi(arr[0])
	if err != nil {
		return 1005
	}

	y, err := strconv.Atoi(arr[1])
	if err != nil {
		return 1005
	}

	if x > 0 && x <= 10 && y > 0 && y <= 10 {
		if operation(text) == "+" {
			total := x + y
			return total
		}
		if operation(text) == "-" {
			total := x - y
			return total
		}
		if operation(text) == "/" {
			total := x / y
			return total
		}
		if operation(text) == "*" {
			total := x * y
			return total
		}
	}
	return 1007
}

// *******
// Функция, проверки арабских и римских чисел в тексте.
func checkArabOrRoman(text string) string {
	arr := arrayLen(text)
	arr1 := arr[0]
	arr2 := arr[1]
	check1 := 0
	check2 := 0
	for i := range arr1 {
		if arr1[i] == 73 || arr1[i] == 86 || arr1[i] == 88 {
			check1 = 1
		} else if arr1[i] == 48 || arr1[i] == 49 || arr1[i] == 50 || arr1[i] == 51 || arr1[i] == 52 || arr1[i] == 53 || arr1[i] == 54 || arr1[i] == 55 || arr1[i] == 56 || arr1[i] == 57 {
			check1 = 3
		} else {
			check1 = 5
		}
	}
	for i := range arr2 {
		if arr2[i] == 73 || arr2[i] == 86 || arr2[i] == 88 {
			check2 = 1
		} else if arr2[i] == 48 || arr2[i] == 49 || arr2[i] == 50 || arr2[i] == 51 || arr2[i] == 52 || arr2[i] == 53 || arr2[i] == 54 || arr2[i] == 55 || arr2[i] == 56 || arr2[i] == 57 {
			check2 = 3
		} else {
			check2 = 5
		}
	}
	if check1 == 1 {
		if check2 == 1 {
			return "roman"
		}
		if check2 == 3 {
			return "osh"
		}
		if check2 == 5 {
			return "osh4"
		}
	}
	if check1 == 3 {
		if check2 == 1 {
			return "osh"
		}
		if check2 == 3 {
			return "arab"
		}
		if check2 == 5 {
			return "osh4"
		}
	}
	if check1 == 5 {
		if check2 == 1 {
			return "osh4"
		}
		if check2 == 3 {
			return "osh4"
		}
		if check2 == 5 {
			return "osh4"
		}
	}
	return "osh4"
}

// *******
// Функция, которая совмещает в себе функции обработки, проверки и проведения математических операций арабских и римских чисел. Функция выдает результат.
func arOperation(text string) string {
	lookOperator := operation(text)
	var total string
	var totalInt int
	var arr []string
	switch lookOperator {
	case "+":
		arr = arrayLen(text)
		check := checkArabOrRoman(text)
		if check == "roman" {
			total = arrOperationRoman(text, arr)
			return total
		}
		if check == "arab" {
			totalInt = arrOperationArab(text, arr)
			totalIntText := strconv.Itoa(totalInt)
			return totalIntText
		}
		if check == "osh" {
			return "1008"
		}
		if check == "osh4" {
			return "1005"
		}
	case "-":
		arr = arrayLen(text)
		check := checkArabOrRoman(text)
		if check == "roman" {
			total = arrOperationRoman(text, arr)
			return total
		}
		if check == "arab" {
			totalInt = arrOperationArab(text, arr)
			totalIntText := strconv.Itoa(totalInt)
			return totalIntText
		}
		if check == "osh" {
			return "1008"
		}
		if check == "osh4" {
			return "1005"
		}
	case "*":
		arr = arrayLen(text)
		check := checkArabOrRoman(text)
		if check == "roman" {
			total = arrOperationRoman(text, arr)
			return total
		}
		if check == "arab" {
			totalInt = arrOperationArab(text, arr)
			totalIntText := strconv.Itoa(totalInt)
			return totalIntText
		}
		if check == "osh" {
			return "1008"
		}
		if check == "osh4" {
			return "1005"
		}
	default:
		arr = arrayLen(text)
		check := checkArabOrRoman(text)
		if check == "roman" {
			total = arrOperationRoman(text, arr)
			return total
		}
		if check == "arab" {
			totalInt = arrOperationArab(text, arr)
			totalIntText := strconv.Itoa(totalInt)
			return totalIntText
		}
		if check == "osh" {
			return "1008"
		}
		if check == "osh4" {
			return "1005"
		}
	}
	return "1005"
}

func main() {
	fmt.Println("Напишите какую операцию надо выполнить (Пример: 3 + 5; VIII / V):")
	enterText := bufio.NewScanner(os.Stdin)
	enterText.Scan()
	text := enterText.Text()

	checkOperator := operation(text)

	if checkOperator == "+" || checkOperator == "-" || checkOperator == "*" || checkOperator == "/" {
		arrayText := arrayLen(text)

		if len(text) <= 2 && len(text) > 1 {
			fmt.Println("Ошибка1002, неправильный формат математической операции. Введите две арабские цифры (от 1 до 10) или две арабские цифры (от I до X) и один оператор (+, -, /, *).")
			return
		}
		if len(arrayText) <= 2 && len(arrayText) > 1 {
			arOperations := arOperation(text)

			if arOperations == "1006" {
				fmt.Println("Ошибка1006, неправильный формат римских цифр. Калькулятор принимает римские цифры от I до X.")
				return
			}
			if arOperations == "1005" {
				fmt.Println("Ошибка1005, неправильный формат математической операции. Введите две арабские цифры (от 1 до 10) или две арабские цифры (от I до X) и один оператор (+, -, /, *).")
				return
			}
			if arOperations == "1008" {
				fmt.Println("Ошибка1008, так как используются одновременно разные системы счисления.")
				return
			}
			if arOperations == "1007" {
				fmt.Println("Ошибка1007, неправильный формат арабских цифр. Калькулятор принимает арабские цифры от 1 до 10.")
				return
			}
			if arOperations == "1009" {
				fmt.Println("Исключение, так как в римской системе нет отрицательных чисел.")
				return
			}
			fmt.Println(arOperations)
			return
		} else {
			fmt.Println("Ошибка1004, неправильный формат математической операции. Введите две арабские цифры (от 1 до 10) или две арабские цифры (от I до X) и один оператор (+, -, /, *).")
			return
		}
	}
	fmt.Println(checkOperator)
	return
}
