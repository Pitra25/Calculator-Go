package main

import (
	"Calculator-Go/src"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello User!! \n To exit, enter exit. \n To view history, enter.")

	for {
		example := readAndSplit()

		if example == "exit" {
			break
		}

		fmt.Println("Result: " + src.Calculation(example))
	}
	fmt.Print("Completion of work...")
}

func readAndSplit() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter example: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
