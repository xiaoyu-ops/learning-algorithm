package main
import (
	"sort";
	"fmt";
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
// 这个方法的精髓就是先排序，然后我们再锚定一个数假设其已经在结果集中
// 然后用双指针法在剩下的数中寻找另外两个数，使得三数之和为0
// 在寻找过程中，我们还要跳过重复的数以避免重复解
// 注意最精髓的部分就是锚定这一步，因为锚定一个数后，问题就简化为在剩下的数
// 中寻找两个数使得它们的和为一个固定值

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}