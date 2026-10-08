package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		if value >= 'a' && value <= 'z' {
			runes[index] = 'a' + (value-'a'+rune(key))%26
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
		if value >= 'a' && value <= 'z' {
			runes[index] = 'a' + (value-'a'-rune(key)+26)%26
		} else {
			runes[index] = value
		}

	}
	return string(runes)
}

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	for {
		for {
			fmt.Println("Enter your mode: ")
			fmt.Println("E - Encrypt Mode")
			fmt.Println("D - Decrypt Mode")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "E" || input == "e" || input == "D" || input == "d" {
				break
			} else {
				fmt.Print("Not a mode! Try again\n\n")
			}
		}
		if input == "E" || input == "e" {
			for {
				fmt.Print("\nEnter your string \nor\nQ to Quit\nor\nR to change mode\n:")
				inString, _ := reader.ReadString('\n')
				inString = strings.TrimSpace(inString)
				if inString == "Q" || inString == "q" {
					fmt.Println("Goodbye!\n")
					return
				}
				if inString == "R" || inString == "r" {
					fmt.Println("Switching Role\n")
					break
				}
				fmt.Print("\nEnter your key\n:")
				key, _ := reader.ReadString('\n')
				key = strings.TrimSpace(key)

				intKey, err := strconv.Atoi(key)

				if err != nil {
					fmt.Println("Invalid Key Given", err)
					return
				}

				out := Encrypt(inString, intKey)
				fmt.Println()
				fmt.Println(out)

			}
		} else if input == "D" || input == "d" {
			for {
				fmt.Print("\nEnter your string \nor\nQ to Quit\nor\nR to change mode\n:")
				inString, _ := reader.ReadString('\n')
				inString = strings.TrimSpace(inString)
				if inString == "Q" || inString == "q" {
					fmt.Println("Goodbye!\n")
					return
				}
				if inString == "R" || inString == "r" {
					fmt.Println("Switching Role\n")
					break
				}
				fmt.Print("\nEnter your key\n:")
				key, _ := reader.ReadString('\n')
				key = strings.TrimSpace(key)

				intKey, err := strconv.Atoi(key)

				if err != nil {
					fmt.Println("Invalid Key Given", err)
					return
				}

				out := Decrypt(inString, intKey)
				fmt.Println()
				fmt.Println(out)
			}
		}
	}

}
