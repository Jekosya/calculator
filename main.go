package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для проверки, является ли строка римским числом
func isRomanNumeral(str string) bool {
	// Регулярное выражение для проверки римских чисел
	romanRegex := regexp.MustCompile(`^(I|II|III|IV|V|VI|VII|VIII|IX|X)$`)
	return romanRegex.MatchString(str)
}

// Функция для выполнения операции сложения
func add(a, b int) int {
	return a + b
}

// Функция для выполнения операции вычитания
func subtract(a, b int) int {
	return a - b
}

// Функция для выполнения операции умножения
func multiply(a, b int) int {
	return a * b
}

// Функция для выполнения операции деления
func divide(a, b int) int {
	return a / b
}

// Основная функция программы
func main() {
	var input string
	fmt.Println("Это калькулятор простых арифметических операций. Введите операцию между числами от 1 до 10 включительно. Между числами и операндом обязательно нужно поставить пробелю  Числа можно вводить в рабском и римском формате, но нельзя проводить операции в разных системах счисления.")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // Считываем строку ввода с консоли
	if err != nil {
		panic(err) // При возникновении ошибки паникуем
	}

	parts := strings.Fields(input) // Разбиваем строку на части
	if len(parts) != 3 {
		panic("Неверный формат ввода. Ожидалось 'число оператор число'.")
	}

	// Преобразуем первый аргумент в число или проверяем, является ли он римским числом
	var num1 int
	if n, err := strconv.Atoi(parts[0]); err == nil && n >= 1 && n <= 10 {
		num1 = n
	} else if isRomanNumeral(parts[0]) {
		num1 = getArabicFromRoman(parts[0])
	} else {
		panic("Первый аргумент должен быть числом от 1 до 10 или римским числом. Или вы проводите операцию в разных системах счисления.")
	}

	// Преобразуем второй аргумент в число или проверяем, является ли он римским числом
	var num2 int
	if n, err := strconv.Atoi(parts[2]); err == nil && n >= 1 && n <= 10 {
		num2 = n
	} else if isRomanNumeral(parts[2]) {
		num2 = getArabicFromRoman(parts[2])
	} else {
		panic("Второй аргумент должен быть числом от 1 до 10 или римским числом. Или вы проводите операцию в разных системах счисления.")
	}

	operator := parts[1] // Получаем оператор из второй части строки
	var result int

	// Проверка на соответствие типов чисел
	if (isRomanNumeral(parts[0]) && !isRomanNumeral(parts[2])) || (!isRomanNumeral(parts[0]) && isRomanNumeral(parts[2])) {
		panic("Калькулятор не может работать с разными типами чисел одновременно.")
	}

	// Выполняем операцию в зависимости от оператора
	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		panic("Неподдерживаемая операция.")
	}

	// Если оба числа римские, выводим результат в римской нотации
	if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) {
		if result <= 0 {
			panic("Результат работы калькулятора с римскими числами не может быть отрицательным или равным нулю.")
		}
		fmt.Printf("Результат: %s\n", getRomanFromArabic(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}

// Функция для преобразования римских чисел в арабские
func getArabicFromRoman(roman string) int {
	// Соответствия римских и арабских чисел
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	var result int
	for i := 0; i < len(roman); i++ {
		if i > 0 && romanNumerals[rune(roman[i])] > romanNumerals[rune(roman[i-1])] {
			result += romanNumerals[rune(roman[i])] - 2*romanNumerals[rune(roman[i-1])]
		} else {
			result += romanNumerals[rune(roman[i])]
		}
	}
	return result
}

// Функция для преобразования арабских чисел в римские
func getRomanFromArabic(num int) string {
	// Соответствия арабских и римских чисел
	romanNumerals := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
		50: "L", 100: "C",
	}

	// Рекурсивно получаем римскую нотацию
	if num <= 10 {
		return romanNumerals[num]
	} else if num < 50 {
		return "X" + getRomanFromArabic(num-10)
	} else if num < 100 {
		return "L" + getRomanFromArabic(num-50)
	} else if num < 500 {
		return "C" + getRomanFromArabic(num-100)
	} else {
		panic("Результат работы калькулятора с римскими числами не может быть больше C.")
	}
}
