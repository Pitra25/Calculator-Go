package src

import (
	"fmt"
	"io"
	"os"
	"time"
)

func saveHistory(example string, result float64) {
	if hasNonZeroDecimalPart(result) {
		writeToFile(fmt.Sprintf("%s = %.2f\n", example, result))
	} else {
		writeToFile(fmt.Sprintf("%s = %d\n", example, int(result)))
	}
}

func writeToFile(text string) {
	fileName := "history.txt"
	// Get current date and time
	now := time.Now()

	data := now.Format("2006-01-02 15:04:05")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
	const lengthSlice int8 = 64

	file, err := os.Open("history.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			data := make([]byte, lengthSlice)
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

func hasNonZeroDecimalPart(num float64) bool {
	return num != float64(int(num))
}
