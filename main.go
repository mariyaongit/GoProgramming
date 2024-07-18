package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func CountVowels(userstring string) int {
	vowels := "aAeEiIoOuU"
	count := 0
	for _, char := range userstring {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string to check the number of vowels: ")
	userString, _ := reader.ReadString('\n')
	userString = strings.TrimSpace(userString) 
	vowelsCount := CountVowels(userString)
	fmt.Println("Number of vowels in the given string:", vowelsCount)
}
