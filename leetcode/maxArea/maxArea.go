package main

import(
	"fmt"
)
// 用双指针，一个指向头，一个指向尾
// 计算面积，更新最大面积
// 然后移动较短的指针，直到两个指针相遇

func maxArea(height []int) int {
    maxArea := 0
	for i,j:=0, len(height) - 1;i < j;{
		area := (j - i) * min(height[i], height[j])
		if area > maxArea{
			maxArea = area
		}
		if height[i] < height[j]{
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func main(){
	height := []int{1,8,6,2,5,4,8,3,7}
	answer := maxArea(height)
	fmt.Printf("the answer is :%v\n", answer)
}