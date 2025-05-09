package utils

import "fmt"

func PromptInt(message string) int {
	var input int
	fmt.Print(message)
	fmt.Scanln(&input)
	return input
}
