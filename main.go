package main

import (
	"Calculator-Go/src"
	"fmt"
)

func main() {
	fmt.Println("Hello User!! \n To exit, enter exit. \n To view history, enter.")

	for {
		example := ""

		fmt.Print("Enter example: ")
		_, err := fmt.Scanln(&example)
		if err != nil {
			fmt.Println(err)
		}

		if example == `exit` {
			break
		}

		fmt.Println("Result: " + src.Calculation(example))
	}
	fmt.Print("Completion of work...")
}
