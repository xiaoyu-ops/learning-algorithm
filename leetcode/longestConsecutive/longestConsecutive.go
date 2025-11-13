package main

import(
	"fmt"
)

func longestConsecutive(nums []int) int {
    check := map[int]bool{}
	for _,num := range nums{
		check[num] = true
	} 
	longest := 0
	longest_now := 0
	for num := range check{
		// 只从序列的起点开始计数
		if !check[num - 1]{
			longest_now = 1
			for check[num + 1]{
				num++
				longest_now++
			}
			if longest_now > longest{
				longest = longest_now
			}
		}
	}
	return longest
}

func main(){
	nums := []int{100,4,200,1,3,2}
	answer := longestConsecutive(nums)
	fmt.Printf("the answer is :%v\n", answer)
}
