package main

import (
	"fmt"
)

func main() {

	var arr = [...]int{6, 5, 4, 3, 2, 1}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}

}
