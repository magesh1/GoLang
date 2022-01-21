package main

import (
	"fmt"
)

func main() {

	
	var arr[3] string
	arr[0] = "Hello"
	arr[1] = "World"
	arr[2] = "!"
	fmt.Println(arr[1])
	

	//shorthand declaration

	arr1 := [3]string{"Hello", "World", "!"}

	for i := 0; i < len(arr1) ; i++ {
		fmt.Println(arr1[i])
	}

	// creating 2d array

	arr2 := [2][2]int{{1,1},{2,2}}

	for i := 0;i < len(arr2);i++ {
		for j := 0;j < len(arr2);j++ {
			fmt.Print(arr2[i][j]," ")
		}
		fmt.Println()
	}

	var arr3[2][2] int
	arr3[0][0] = 1
	arr3[0][1] = 2
	arr3[1][0] = 3
	arr3[1][1] = 4

	for i := 0;i < len(arr3);i++ {
		for j := 0;j < len(arr3);j++ {
			fmt.Print(arr3[i][j]," ")
		}
		fmt.Println()
	}

	// creating dynamic arr3
	var arr4 = [...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(len(arr4))


}
