package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Function CountVowels count  the number of vowels in the passed string and return the count,
func CountVowels(userstring string) int {
	vowels := "aAeEiIoOuU"  //string containing vowels
	count := 0  // count is intialized to 0
	for _, char := range userstring {
		if strings.ContainsRune(vowels, char) { // check if the specific character is present in a string.
			count++
		}
	}
	return count
}

func main() {
	//Read user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string to check the number of vowels: ")
	userString, _ := reader.ReadString('\n')
	userString = strings.TrimSpace(userString) //remove leading and trailing whitespace
	vowelsCount := CountVowels(userString) // call CountVowels to count vowels
	fmt.Println("Number of vowels in the given string:", vowelsCount)
}
