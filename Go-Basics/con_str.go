package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "WelcomE"
	str2 := "Basics of Golang"
	str3 := "HELLO! "

	fmt.Println("Strings (before):")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)
	fmt.Println("String 3:", str3)

	res1 := strings.ToLower(str1)
	res2 := strings.ToLower(str2)
	res3 := strings.ToLower(str3)

	fmt.Println("\nStrings (after):")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3:", res3)

}
