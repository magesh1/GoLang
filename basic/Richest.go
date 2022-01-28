# https://leetcode.com/problems/richest-customer-wealth/

package main

import "fmt"

func main() {
  accounts := [][]int{{1,2,3},{3,2,1}}
  maximumWealth(accounts)
  
}

func maximumWealth(accounts [][]int) {
    
    res := make([]int, len(accounts))

	for i := 0; i < len(accounts); i++ {
		var sum int
		for _, j := range accounts[i] {
			sum += j
		}
		res[i] = sum
	}

	check := res[0]

	for i := 1; i < len(res); i++ {
		if check < res[i] {
			check = res[i]
		}
	}
    
  fmt.Println(check)

    
}
