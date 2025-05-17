package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello User!!")
	for {
		var (
			arrayNumber      []float64
			char             string
			historyLineCount int = 5
		)

		fmt.Print("Enter numbers separated by spaces: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		for _, element := range stringToFloat(input) {
			arrayNumber = append(arrayNumber, element)
		}

		fmt.Print("Enter the sign for calculation: ")
		fmt.Scan(&char)

		if arrayNumber[0] == 0 {
			break
		}

		switch char {
		case "+":
			fmt.Println(sum(arrayNumber...))
		case "-":
			fmt.Println(subtraction(arrayNumber...))
		case "*":
			fmt.Println(multiplication(arrayNumber...))
		case "/":
			fmt.Println(division(arrayNumber...))
		case "h":
			readingFromFile(historyLineCount)
		}
	}
	fmt.Print("Completion of work...")
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
