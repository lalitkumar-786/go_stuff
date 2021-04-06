package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter a number ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	text := input.Text()
	i, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Error")
	}
	if len(text) > 2 {
		fmt.Println("Please enter a single number")
	} else {
		if i > 0 && i <= 10 {
			fmt.Println("Yes..entered number is right")
		}
	}

}
