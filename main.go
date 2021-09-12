package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Проверяет содержится ли элемент в массиве
func Contains(elements []rune, element rune) bool {
	for _, n := range elements {
		if element == n {
			return true
		}
	}
	return false
}

func Counter(s string) string {
	var (
		// Массив шаблонов, который содержит уже "подсчитанные элементы"
		patterns []rune

		inPatterns  bool
		shortString string
	)

	for _, v := range s {

		// Проверяем находится ли шаблон в общем массиве ранее использованных шаблонов
		inPatterns = Contains(patterns, v)

		// Если шаблон не содержится в массиве
		if !inPatterns {
			// Добавляем шаблон в массив
			patterns = append(patterns, v)

			// Подсчитываем количество вхождений шаблона в строку
			pattern := regexp.MustCompile(string(v))
			matches := pattern.FindAllStringIndex(s, -1)

			shortString += string(v) + strconv.Itoa(len(matches))
		}
	}
	return shortString
}

func main() {
	fmt.Printf("Output: %v", Counter("aaabbbccccc"))
}
