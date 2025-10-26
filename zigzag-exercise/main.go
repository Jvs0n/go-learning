package main

import (
	"fmt"
	"unicode"
)

func main() {
	var name = ""
	fmt.Println("Converting your name into zigzag-case")
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	var result []rune
	for index, character := range name {
		if index%2 == 0 {
			result = append(result, unicode.ToLower(character))
		} else {
			result = append(result, unicode.ToUpper(character))
		}
	}
	zigzag := string(result)
	print(zigzag)

}

//day one: Basics
