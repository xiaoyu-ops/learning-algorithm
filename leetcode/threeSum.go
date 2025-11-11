package main
import (
	"sort";
	"fmt";
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	check := map[int]int{}
	
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			value = -(nums[i] + nums[j])
}


func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	fmt.Println(nums)
}