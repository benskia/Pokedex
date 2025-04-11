package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	splitTexts := strings.Split(text, " ")
	cleanTexts := []string{}

	for _, s := range splitTexts {
		trimmedStr := strings.TrimSpace(s)
		if trimmedStr != "" {
			cleanTexts = append(cleanTexts, trimmedStr)
		}
	}

	return cleanTexts
}
