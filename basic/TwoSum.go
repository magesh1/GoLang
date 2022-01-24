// nums = [2,7,11,15], target = 9

package main

import "fmt"

func main() {

	var num = []int{3, 2, 4}
	var target = 6

	twoSum(num, target)

}

func twoSum(num []int, target int) {


	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {
			if num[i]+num[j] == target {
				fmt.Println(i, j)
				break
			}
		}
	}

	

	


}
