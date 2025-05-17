package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello User!!")
	for {
		var (
			num1, num2 float64
			char       string
		)
		fmt.Print("Enter numbers separated by spaces: ")
		fmt.Scanln(&num1, &num2)
		if num1 == 0 && num2 == 0 {
			break
		}
		fmt.Print("Enter the sign for calculation: ")
		fmt.Scan(&char)

		switch char {
		case "+":
			{
				result := num1 + num2
				if hasNonZeroDecimalPart(result) {
					fmt.Printf("%.2f + %.2f = %.3f\n", num1, num2, result)
				} else {
					fmt.Printf("%d + %d = %d\n", int(num1), int(num2), int(result))
				}
			}
		case "-":
			{
				result := num1 - num2
				if hasNonZeroDecimalPart(result) {
					fmt.Printf("%.2f - %.2f = %.3f\n", num1, num2, result)
				} else {
					fmt.Printf("%d - %d = %d\n", int(num1), int(num2), int(result))
				}
			}
		case "*":
			{
				result := num1 * num2
				if hasNonZeroDecimalPart(result) {
					fmt.Printf("%.2f * %.2f = %.3f\n", num1, num2, result)
				} else {
					fmt.Printf("%d * %d = %d\n", int(num1), int(num2), int(result))
				}
			}
		case "/":
			{
				if num2 == 0 {
					fmt.Println("Invalid input")
					break
				}
				result := num1 / num2
				if hasNonZeroDecimalPart(result) {
					fmt.Printf("%.2f / %.2f = %.3f\n", num1, num2, result)
				} else {
					fmt.Printf("%d / %d = %d\n", int(num1), int(num2), int(result))
				}
			}
		}
	}
	fmt.Print("Completion of work...")
}

func hasNonZeroDecimalPart(num float64) bool {
	return num != float64(int(num))
}
