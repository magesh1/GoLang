# https://leetcode.com/problems/concatenation-of-array/

package main

import "fmt"

func main() {
  nums := []int{1,2,1}
  getConcatenation(nums)
  
}



func getConcatenation(nums []int) {
    
    nums = append(nums,nums...)
  fmt.Println(nums)
}
