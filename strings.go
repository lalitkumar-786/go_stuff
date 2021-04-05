// String operations basic

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Yes. I 've uploaded this code"

	fmt.Println("The original string is :", s)

	// caculate len
	fmt.Println(len(s))

	//partition it from index 2-4
	fmt.Println(s[1:7])

	// // split it
	text := strings.Split(s, " ")
	fmt.Println(text)

	// //to iterate over it
	for i := 0; i < len(text); i++ {
		fmt.Print(text[i], "	")
	}
	fmt.Println()
	//better way to iterate string in go
	// NOTE the below val will be ASCII code. To convert it to string, using string(val)
	for key, val := range s {
		fmt.Println("At index ", key, "the char is ", string(val))
	}
}
