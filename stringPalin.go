package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func isPalindromeLong(s string) {

	// Prefer RuneCountInString as len(s) --> returns the bytes not the chars
	l := utf8.RuneCountInString(s)
	count := 0
	flag := 0
	for i := 0; i < l/2; i++ {
		if s[i] == s[l-i-1] {
			count++
		} else {
			fmt.Println("Sorry.. Strings are not Palindrome")
			flag = 1
			break
		}
	}

	if flag == 0 {
		fmt.Println("Given string is Palindrome")
	}
}

func isPalindromeRev(s string) {

	res := ""

	for _, val := range s {
		res = string(val) + res
	}

	if s == res {
		fmt.Println("Yes.. It's a Palindrome")
	} else {
		fmt.Println("No..Given string is not Palindrome")
	}
}

func main() {

	fmt.Println("Enter a string:-")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	text := input.Text()

	// check palindrome using the basic brute force
	isPalindromeLong(text)

	//Check palindrom using reverse string
	isPalindromeRev(text)

}
