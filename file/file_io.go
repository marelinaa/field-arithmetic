package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadPolynom(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}
	return "", fmt.Errorf("file %s is empty", filename)
}

func ReadInput(filename string) (operation, poly1, poly2 string, k int, err error) {
	// Открываем файл
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", "", "", 0, err
	}

	// Читаем строку и удаляем пробелы
	line := strings.TrimSpace(string(content))

	// Определяем оператор
	var operator string
	switch {
	case strings.Contains(line, "+"):
		operator = "+"
	case strings.Contains(line, "*"):
		operator = "*"
	case strings.Contains(line, "/"):
		operator = "/"
	case strings.Contains(line, "^"):
		operator = "^"
	default:
		return "", "", "", 0, errors.New("invalid input format")
	}

	// Разбираем строку в зависимости от оператора
	switch operator {
	case "+", "*", "/":
		parts := strings.Split(line, operator)
		if len(parts) != 2 {
			return "", "", "", 0, errors.New("invalid format for binary operation")
		}
		poly1 = strings.TrimSpace(parts[0])
		poly2 = strings.TrimSpace(parts[1])
		if poly1 == "" || poly2 == "" {
			return "", "", "", 0, errors.New("one of the polynoms is empty")
		}
		return operator, poly1, poly2, 0, nil
	case "^":
		parts := strings.Split(line, "^")
		if len(parts) != 2 {
			return "", "", "", 0, errors.New("invalid format for exponentiation")
		}
		poly1 = strings.TrimSpace(parts[0])
		exp := strings.TrimSpace(parts[1])

		// Преобразуем степень в int
		k, err = strconv.Atoi(exp)
		if err != nil {
			return "", "", "", 0, errors.New("invalid exponent value")
		}
		return operator, poly1, "", k, nil
	}
	fmt.Printf("poly: %s; %s; %s; %d\n", operator, poly1, poly2, k)

	return "", "", "", 0, errors.New("unexpected error")
}

func WriteOutput(filename, result string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(result + "\n")
	return err
}
