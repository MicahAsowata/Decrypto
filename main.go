package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFromTerminal(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	str, err := reader.ReadString('\n')
	return str, err
}

func main() {
	word, _ := readFromTerminal("Enter a cipher text")
	number, _ := readFromTerminal("Enter a the key")
	trimmedNumber := strings.Trim(number, "\n")
	intifiedNumber, _ := strconv.Atoi(trimmedNumber)

	fmt.Println("The Plain Text is: ")
	decrypt(word, intifiedNumber)
}
