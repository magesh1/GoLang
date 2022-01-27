package main

import "fmt"

func main() {
	fmt.Println("paliwokrig")

	str1 := "madama"
	
	if findPali(str1) {
		fmt.Println("is palindrome")
	} else {
		fmt.Println("is not palindrome")
	}



}


func findPali(str string) bool {

	res := ""

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	if res == str {
		return true
	}

	return false
}
