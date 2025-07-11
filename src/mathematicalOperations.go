package src

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculation(example string) string {
	const (
		openingParenthesis string = "("
		closingParenthesis string = ")"
		errorResult        string = "Ошибка вода"
	)

	var (
		result                     float64
		char                       string
		numberFloat1, numberFloat2 float64
		historyLineCount           uint8 = 5
		bracket                          = true
		indexNum1                        = 0
		indexNum2                        = 2
		indexSymbol                      = 1
	)

	if strings.Contains(example, openingParenthesis) &&
		!strings.Contains(example, closingParenthesis) {
		return errorResult
	}

	items := []rune(example)

	//nolint:all
	if !bracket {
		gfg
	}

	for index := 0; index <= len(items); {

		numberFloat1, _ = strconv.ParseFloat(string(items[indexNum1]), 64)
		char = string(items[indexSymbol])
		numberFloat2, _ = strconv.ParseFloat(string(items[indexNum2]), 64)

		switch char {
		case `+`:
			result = sum(numberFloat1, numberFloat2)
		case `-`:
			result = subtraction(numberFloat1, numberFloat2)
		case `*`:
			result = multiplication(numberFloat1, numberFloat2)
		case `/`:
			result = division(numberFloat1, numberFloat2)
		case `h`:
			readingFromFile(historyLineCount)
		}

		if indexNum2 == len(items)-1 {
			break
		}

		indexNum1 = +4
		indexNum2 = +2
		indexSymbol = +1
	}

	saveHistory(example, result)
	return fmt.Sprintf("%f", result)
}

func sum(number1, number2 float64) float64 {
	return number1 + number2
}
func subtraction(number1, number2 float64) float64 {
	return number1 - number2
}
func multiplication(number1, number2 float64) float64 {
	return number1 * number2
}
func division(number1, number2 float64) float64 {
	if number1 == 0 {
		fmt.Println("Invalid input")
		return 0.0
	}

	return number1 / number2
}

func stringToFloat(num string) []float64 {
	resultNum := strings.Split(num, " ")
	var result []float64

	// Преобразуем каждый элемент
	for _, str := range resultNum {
		number, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error while converting:", err)
			return result
		}

		result = append(result, number)
	}

	return result
}
