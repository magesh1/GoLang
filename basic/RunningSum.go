# https://leetcode.com/problems/running-sum-of-1d-array/

package main
import "fmt"

func main() {
  nums := []int{1,2,3,4} 
  findSum(nums)
}

func findSum(nums []int) {

for i := 1;i < len(nums);i++ {
        nums[i] = nums[i-1]+nums[i]
}
    
    
  fmt.Println(nums)
  
}
