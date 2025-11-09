package main

import "fmt"
// 最简单的暴力解法
func twoSum(nums []int, target int) []int {
	for i:= 0; i < len(nums); i++ {
		for j:= i + 1;j< len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
// 哈希表解法
func twoSumHash(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, num := range nums {
		if j, ok := hashMap[target - num];ok{
			return []int{j, i}
		} else {
			hashMap[num] = i
		}
	}
	return nil
}	




func main() {
	var answer []int
	nums := []int{2,7,11,15}
	target := 9
	answer = twoSum(nums, target)
	fmt.Printf("the answer of normal :%d",answer)
	answer = twoSumHash(nums, target)
	fmt.Printf("the answer of hash :%d",answer)
	// 注意fmt.Printf的格式化输出，
	// fmt.println是直接输出变量内容的
}