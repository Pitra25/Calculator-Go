package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello User!!")
	fmt.Println("To exit, enter exit. \n To view history, enter.")

	for {
		var (
			example string
			result  string
		)

		fmt.Print("Enter example: ")
		fmt.Scanln(&example)

		if example == `exit` {
			break
		}

		result = calculation(example)
		fmt.Println("Result: " + result)
	}
	fmt.Print("Completion of work...")
}
