package main

import(
	"fmt"
)

func moveZeroes(nums []int)  {
    for i, j :=0,0;j<len(nums);j++{
		if nums[j] != 0{
			nums[i] = nums[j]
			if i != j{
				nums[j] = 0
			}
			i++
		}
	}
}

func main(){
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Printf("the answer is :%v\n", nums)
}