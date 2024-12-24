package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Функция для распаковки строки, принимает на вход строку, возвращает ошибку либо строку
func funcString(inputString string) (string, error) {
	// Возвращаем пустую строку если на вход была подана пустая строка
	if inputString == "" {
		return "", nil
	}
	// Переменная для хранения результата
	var result string
	// Переменная для хранения последнего символа
	var lastRune rune

	// Цикл для прохода по всем символам строки
	for _, curRune := range inputString {
		// Проверка является ли символ числом иначе сохраняем символ без изменений
		if curRune >= '0' && curRune <= '9' {
			// Проверка на случай если цифра была введена а символ строки еще не введен
			if lastRune == 0 {
				return "", errors.New("Введена некорректная строка!")
			}
			// Попытка преобразования символа в числа, если не выходит возвращаем ошибку
			count, err := strconv.Atoi(string(curRune))
			if err != nil {
				return "", err
			}
			// Конкатенация к строке результата предыдущего дублированного символа уменьшанного на 1
			result += strings.Repeat(string(lastRune), count-1)
		} else {
			// Конкатенция к строке результата текущего символа
			result += string(curRune)
		}
		// Сохранение текущего символа в переменную предыдущего символа
		lastRune = curRune
	}
	return result, nil
}

func main() {
	// Вывод сообщения о вводе строки
	fmt.Print("Введите строку для распаковки: ")
	// Инициализация переменной для хранения строки
	var inputString string
	// Считывание строки
	fmt.Scan(&inputString)
	// Применение функции распаковки строки
	result, err := funcString(inputString)
	// Вывод ошибки или строки в зависимости от результата функции
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Вывод: %q\n", result)
	}
}
