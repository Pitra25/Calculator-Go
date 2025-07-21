package src

import (
	"Calculator-Go/src/save"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Calculation(example string) string {
	var (
		result float64
		//historyLineCount string = "5"
	)

	parts := strings.Fields(example)
	if len(parts) != 3 {
		log.Fatalf("error: Invalid input format. Use format: number operator number")
		return ""
	}

	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		log.Fatalf("error parse float: %v", err)
	}

	char := parts[1]

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		log.Fatalf("error parse float: %v", err)
	}

	switch char {
	case `+`:
		result = sum(num1, num2)
	case `-`:
		result = subtraction(num1, num2)
	case `*`:
		result = multiplication(num1, num2)
	case `/`:
		result = division(num1, num2)
	case `h`:
		return "history" //GetHistory(historyLineCount)
	default:
		return "Вы ввели неизвестное что то :)"
	}

	errS := save.SaveHistory(example, result)
	if errS != nil {
		return fmt.Sprintf("%.2f	|	Результат не сохранен: %v", result, errS)
	}
	return fmt.Sprintf("%.2f", result)
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

func stringToFloat(num string) []float32 {
	resultNum := strings.Split(num, " ")
	var result []float32

	// Преобразуем каждый элемент
	for _, str := range resultNum {
		number, err := strconv.ParseFloat(str, 32)
		if err != nil {
			fmt.Println("Error while converting:", err)
			return result
		}

		result = append(result, float32(number))
	}

	return result
}
