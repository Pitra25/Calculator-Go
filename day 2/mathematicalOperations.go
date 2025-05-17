package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func sum(numbers ...float64) float64 {
	result := 0.0
	for _, number := range numbers {
		result += number
	}
	saveHistory(result, "+", numbers...)
	return result
}
func subtraction(numbers ...float64) float64 {
	result := 0.0
	i := 1
	for j := i + 1; j < len(numbers); j++ {
		result = numbers[i] - numbers[j]
	}
	saveHistory(result, "-", numbers...)
	return result
}
func multiplication(numbers ...float64) float64 {
	result := 0.0
	i := 1
	for j := i + 1; j < len(numbers); j++ {
		result = numbers[i] * numbers[j]
	}
	saveHistory(result, "*", numbers...)
	return result
}
func division(numbers ...float64) float64 {
	if numbers[0] == 0 {
		fmt.Println("Invalid input")
		return 0.0
	}

	result := 0.0
	i := 1
	for j := i + 1; j < len(numbers); j++ {
		result = numbers[i] / numbers[j]
	}
	saveHistory(result, "/", numbers...)
	return result
}

func saveHistory(result float64, sign string, numbers ...float64) {
	strResult := ""
	for index, number := range numbers {
		if hasNonZeroDecimalPart(result) {
			strResult += fmt.Sprintf("%.2f ", number)
		} else {
			strResult += fmt.Sprintf("%d ", int(number))
		}
		if index < len(numbers)-1 {
			strResult += fmt.Sprintf("%s ", sign)
		}
	}
	writeToFile(fmt.Sprintf("%s= %.2f\n", strResult, result))
}

func hasNonZeroDecimalPart(num float64) bool {
	return num != float64(int(num))
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
func readingFromFile(index int) {
	file, err := os.Open("history.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for i := 0; i <= index; i++ {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
