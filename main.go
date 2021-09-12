package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

// Сортирует строку в алфавитном порядке
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// "Сворачивает строку"
func Counter(s string) string {
	var (
		// Массив шаблонов, который содержит уже "подсчитанные элементы"
		patterns []rune

		inPatterns  bool
		shortString string
	)

	// Сортируем строку
	s = SortString(s)

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

			// Формируем строку для вывода
			shortString += string(v) + strconv.Itoa(len(matches))
		}
	}
	return shortString
}

func main() {
	fmt.Printf("Output: %v", Counter("zzzzcccUUUuu"))
}
