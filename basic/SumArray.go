package main

import (
	"fmt"
)

func main() {

	var arr = []int{1,2,3,4,5}
	var res = sumArr(arr)

	fmt.Println(res)


}

func sumArr(arr[] int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}



