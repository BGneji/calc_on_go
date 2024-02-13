package main

import (
	"bufio"
	"fmt"

	"os"
	"slices"
	"strconv"
	"strings"
)

func romanToInt(s string) int {

	sum := 0

	if len(s) == 0 {
		return sum
	}

	m := make(map[string]int)

	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	ls := len(s)
	for i := 0; i < ls; i++ {

		a := m[string(s[i])]
		b := 0

		if i+1 < ls {
			b = m[string(s[i+1])]
		}

		if a < b {
			sum += b - a
			i++
		} else {
			sum += a
		}

	}

	// fmt.Println(sum)
	return sum
}

var romanNumerals = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	roman := ""
	for value, numeral := range romanNumerals {
		for num >= value {
			roman += numeral
			num -= value
		}
	}
	return roman
}

func main1() {
	number := 19
	romanNumber := intToRoman(number)
	fmt.Println(romanNumber, '1') // вывод: "XIX"
}

// проверка на не цифры
func str_data(data []string, rim_bool interface{}) bool {
	if rim_bool == true {
		return true
	}
	for _, a := range data {
		intNumber, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Выдача1 паники, так как строка не является математической операцией.")
			return false
			fmt.Println(intNumber)
		} else {
			return true
		}
	}
	return false
}

// проверка на римские цифры
func stringInSlice(a string, c string, rim []string) interface{} {
	found_a := slices.Contains(rim, a)
	found_c := slices.Contains(rim, c)
	// fmt.Println(found_c)
	if found_a == true && found_c == true {
		return true
	}
	if (found_a == true && found_c == false) || (found_a == false && found_c == true) {
		return "num_and_rim"
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите пример через пробел: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSuffix(input, "\n") // Удаление символа новой строки

	data := strings.Fields(input)
	// fmt.Println("Список элементов data:", len(data))
	rim := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	rim_bool := stringInSlice(data[0], data[2], rim)
	// fmt.Println("rim_bool", rim_bool)
	if len(data) == 3 && rim_bool == "num_and_rim" {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления. 1 + I")
		return
	}

	// fmt.Println("rim_bool:", rim_bool)
	str_data := str_data(data, rim_bool)
	// fmt.Println("str_data:", str_data)
	if str_data == false {
		return
	}

	if len(data) < 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		return
	}
	if len(data) > 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	if len(data) == 3 && rim_bool == false {
		// fmt.Println("Цифры")
		num_1, err := strconv.ParseInt(data[0], 10, 64)
		num_2, err := strconv.ParseInt(data[2], 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if (num_1 >= 1 && num_1 <= 10) && (num_2 >= 1 && num_2 <= 10) {
			// fmt.Println((num_1 >= 1 && num_1 <= 10))
			// fmt.Println((num_2 >= 1 && num_2 <= 10))
			// fmt.Println("Решаем примеры")
			if data[1] == "+" {
				fmt.Println(num_1 + num_2)
			}
			if data[1] == "-" {
				fmt.Println(num_1 - num_2)
			}
			if data[1] == "*" {
				fmt.Println(num_1 * num_2)
			}
			if data[1] == "/" {
				fmt.Println(num_1 / num_2)
			}
			return
		}
		fmt.Println("Калькулятор должен принимать на вход числа от 1 до 10 включительно,")

	}

	if len(data) == 3 && rim_bool == true && str_data == true {
		// fmt.Println("Римские цифры")
		if data[1] == "+" {
			q := (romanToInt(data[0]) + romanToInt(data[2]))
			romanNumber := intToRoman(q)
			fmt.Println(romanNumber)
		}
		if data[1] == "-" {
			q := (romanToInt(data[0]) - romanToInt(data[2]))
			if q < 0 {
				fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
				return
			}
			romanNumber := intToRoman(q)
			fmt.Println(romanNumber)
		}
		if data[1] == "*" {
			q := (romanToInt(data[0]) * romanToInt(data[2]))
			romanNumber := intToRoman(q)
			fmt.Println(romanNumber)
		}
		if data[1] == "/" {
			q := (romanToInt(data[0]) / romanToInt(data[2]))
			romanNumber := intToRoman(q)
			fmt.Println(romanNumber)
		}
	}

}
