package main

import (
	"fmt"
	
)

func main() {

	fmt.Println("check the character type")

	var ch string
	fmt.Scanln(&ch)

	if ch >= "a" && ch <= "z" {
		fmt.Println("the character is lowercase")
	} else if ch >= "A" && ch <= "Z" {
		fmt.Println("the character is uppercase")
	}  else if ch >= "0" && ch <= "9" {
		fmt.Println("the character is digit")
	} else {
		fmt.Println("the character is special character")
	}

}
