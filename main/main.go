package main

import (
	"fmt"
	"os"

	"github.com/marelinaa/field-arithmetic/file"
	"github.com/marelinaa/field-arithmetic/operations"
)

const (
	polynomFile = "polynom.txt"
	inputFile   = "input.txt"
	outputFile  = "output.txt"
)

func main() {

	// Чтение неприводимого многочлена
	polynom, err := file.ReadPolynom(polynomFile)
	if err != nil {
		fmt.Println("Error reading polynom file:", err)
		os.Exit(1)
	}

	fmt.Println(polynom)
	// Чтение входных данных
	operation, poly1, poly2, k, err := file.ReadInput(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Выполнение операции
	var result string
	switch operation {
	case "+":
		result = operations.AddPolynomials(poly1, poly2)
	case "*":
		result = operations.MultiplyPolynomials(poly1, poly2)
	case "/":
		fmt.Println("вход в функцию")
		q, r := operations.DividePolynomials(poly1, poly2)
		result = q
		result += r
		fmt.Println("вывод рез-та")
		if err != nil {
			fmt.Println("Error during division:", err)
			os.Exit(1)
		}
	case "^":
		result = operations.PowerPolynomial(poly1, k)
	// case "inv":
	// 	result, err = operations.InversePolynomial(poly1, polynom)
	// 	if err != nil {
	// 		fmt.Println("Error finding inverse:", err)
	// 		os.Exit(1)
	// 	}
	default:
		fmt.Println("Unknown operation:", operation)
		os.Exit(1)
	}

	// Запись результата
	err = file.WriteOutput(outputFile, result)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Operation completed successfully. Result saved to", outputFile)
}
