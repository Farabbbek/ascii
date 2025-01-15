package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Banner() (map[rune][]string, error) {
	// Открывает файл для чтения. Если файл не существует или недоступен, возвращается ошибка.

	file, err = os.Open("standard.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close() //Закрывает файл автоматически после завершения функции, чтобы освободить ресурсы.

	banner := make(map[rune][]string) //Создает мапу, где ключ — это символ
	scanner := bufio.NewScanner(file) //построчно read a file

	var currentRune rune = ' '
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentRune != 0 {
				banner[currentRune] = lines
			}
			currentRune++
			lines = []string{}
		} else {
			lines = append(lines, line)
		}
	}
	if currentRune != 0 {
		banner[currentRune] = lines
	}
	return banner, nil
}

func Ascii(input string, banner map[rune][]string) string {
	var result strings.Builder
	lines := make([]string, 8)

	for _, char := range input {
		if char == '\n' {
			for i := 0; i < 8; i++ {
				lines[i] += "$"
			}
			result.WriteString(strings.Join(lines, "\n") + "\n\n")
			lines = make([]string, 8)
			continue
		}

		asciiArt, exists := banner[char]
		if !exists {
			asciiArt = banner[' ']
		}

		for i := 0; i < 8; i++ {
			lines[i] += asciiArt[i]
		}
	}

	for i := 0; i < 8; i++ {
		lines[i] += "$"
	}
	result.WriteString(strings.Join(lines, "\n"))

	return result.String()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <name>") //Она проверяет, передан ли аргумент командной строки.
		return
	}

	input := os.Args[1]

	//Создаем массив из 8 строк для хранения строк ASCII-арта
	banner, err := Banner()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	output := Ascii(input, banner)
	fmt.Println(output) //Выводим результат на экран

}
