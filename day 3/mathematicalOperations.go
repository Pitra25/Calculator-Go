package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func calculation(example string) string {
	var (
		result                     float64
		char                       string
		numberFloat1, numberFloat2 float64
		arraySymbol                []string
		historyLineCount           uint8 = 5
	)
	//arraySymbol = append(arraySymbol, strings.Split(example, " ")...)
	j := 2
	s := 1
	for i := 0; i <= len(example); i++ {
		numberFloat1, _ = strconv.ParseFloat(arraySymbol[i], 64)
		char = arraySymbol[s]
		numberFloat2, _ = strconv.ParseFloat(arraySymbol[j], 64)

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

func hasNonZeroDecimalPart(num float64) bool {
	return num != float64(int(num))
}

func saveHistory(example string, result float64) {
	if hasNonZeroDecimalPart(result) {
		writeToFile(fmt.Sprintf("%s= %.2f\n", example, result))
	} else {
		writeToFile(fmt.Sprintf("%s= %d\n", example, int(result)))
	}
}
func writeToFile(text string) {
	fileName := "history.txt"
	// Get current date and time
	now := time.Now()

	data := now.Format("2006-01-02 15:04:05")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	_, err = file.WriteString(data + " | " + text)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Done.")
}
func readingFromFile(index uint8) {
	file, err := os.Open("history.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			data := make([]byte, 64)
			var i uint8
			for i = 0; i <= index; i++ {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}
				fmt.Print(string(data[:n]))
			}
		} else {
			fmt.Println("Error reading file")
		}
	}(file)
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
