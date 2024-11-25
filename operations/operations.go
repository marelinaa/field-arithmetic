package operations

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AddPolynomials(poly1, poly2 string) string {
	maxLen := max(utf8.RuneCountInString(poly1), utf8.RuneCountInString(poly2))

	poly1 = poly1 + strings.Repeat("0", maxLen-len(poly1)) + poly1
	poly2 = poly2 + strings.Repeat("0", maxLen-len(poly2))

	result := make([]byte, maxLen)
	for i := 0; i < maxLen; i++ {
		// XOR для каждого разряда
		if poly1[i] == poly2[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}

// Функция для получения i-го бита полинома (по индексу)
func getBit(poly string, index int) int {
	if index < 0 || index >= len(poly) {
		return 0
	}
	return int(poly[len(poly)-1-index] - '0')
}

// MultiplyPolynomials - функция умножения полиномов в поле 2
func MultiplyPolynomials(poly1, poly2 string) string {
	// Определяем длины полиномов
	len1 := len(poly1)
	len2 := len(poly2)

	// Результирующий массив для коэффициентов произведения
	result := make([]int, len1+len2-1)

	// Перемножаем каждый коэффициент
	for i := 0; i < len1; i++ {
		if poly1[len1-1-i] == '1' { // Проверяем коэффициент poly1
			for j := 0; j < len2; j++ {
				if poly2[len2-1-j] == '1' { // Проверяем коэффициент poly2
					result[i+j] ^= 1 // XOR для добавления монома
				}
			}
		}
	}

	// Преобразуем результат в строку
	var builder strings.Builder
	for i := len(result) - 1; i >= 0; i-- {
		builder.WriteByte(byte(result[i] + '0'))
	}

	return builder.String()
}

// Функция XOR для двух строковых полиномов
func XORStrings(p1, p2 string) string {
	maxLen := len(p1)
	if len(p2) > maxLen {
		maxLen = len(p2)
	}

	// Дополняем полиномы до одинаковой длины
	p1 = p1 + strings.Repeat("0", maxLen-len(p1))
	p2 = p2 + strings.Repeat("0", maxLen-len(p2))

	var result strings.Builder
	for i := 0; i < maxLen; i++ {
		if p1[i] != p2[i] {
			result.WriteByte('1')
		} else {
			result.WriteByte('0')
		}
	}

	return strings.TrimLeft(result.String(), "0")
}

// DividePolynomials - функция деления полиномов в поле 2
func DividePolynomials(dividend, divisor string) (q, r string) {
	if len(dividend) < len(divisor) {
		return "0", dividend // cтепень делимого < степени делителя
	}

	q = ""
	r = dividend
	fmt.Println("начало работы")

	for len(r) >= len(divisor) {
		// степень частного
		degreeDiff := len(r) - len(divisor)
		currentQuotient := strings.Repeat("0", degreeDiff) + "1"

		// Добавляем текущий коэффициент в частное
		q = XORStrings(q, currentQuotient)

		// Вычитаем делитель, сдвинутый на степень (XOR)
		shiftedDivisor := divisor + strings.Repeat("0", degreeDiff)
		r = XORStrings(r, shiftedDivisor)
	}

	fmt.Println("конец цикла")
	// Возвращаем частное и остаток
	return strings.TrimLeft(q, "0"), r
}

func PowerPolynomial(poly string, k int) string {
	result := "1"
	base := poly

	for k > 0 {
		if k%2 == 1 { // Если степень нечетная, умножаем на base
			result = MultiplyPolynomials(result, base)
		}
		base = MultiplyPolynomials(base, base)
		k /= 2
	}

	return result
}
