package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Encrypt(input string, key int) string {
	runes := []rune(input)
	key = key % 26

	for index, value := range runes {
		if value >= 'A' && value <= 'Z' {
			runes[index] = 'A' + (value-'A'+rune(key))%26
		} else {
			runes[index] = value
		}
	}

	return string(runes)
}

func Decrypt(input string, key int) string {
	runes := []rune(input)
	key = ((key % 26) + 26) % 26

	for index, value := range runes {
		if value >= 'A' && value <= 'Z' {
			runes[index] = 'A' + (value-'A'-rune(key)+26)%26
		} else {
			runes[index] = value
		}
	}

	return string(runes)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		for {
			fmt.Println("Enter your mode: ")
			fmt.Println("E - Encrypt Mode")
			fmt.Println("D - Decrypt Mode")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "E" || input == "D" {

			} else {
				fmt.Print("Not a mode! Try again\n\n")
			}
		}
	}

}
