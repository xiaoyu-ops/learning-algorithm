package main
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
import (
	"fmt"
)

// “滑动窗口”解法的简单思路：

// 您可以把“窗口”想象成您用两个手指（left 和 right）在字符串上框出的一段。

// “尺蠖”前进：

// left 手指（左边界）先不动，right 手指（右边界）一个一个向右移动，不断扩大窗口。

// 记录足迹：

// right 手指每经过一个新字符，就用一个 map 记下这个字符最后出现的位置（它的索引）。

// 遇到重复：

// right 手指突然遇到了一个在 map 里已经有了，并且位置在 left 手指或其右边的字符。

// 这意味着窗口内出现了重复！

// 窗口收缩：

// 此时，right 手指停下。left 手指（左边界）必须“跳”到那个重复字符上一次出现位置的“下一个”位置。

// （例如，窗口是 [a, b, c, a]，left 在0，right 在3。a 重复了。left 必须跳到索引 0+1=1 的位置，窗口变为 [b, c, a]。）

// 更新答案：

// 在每一步（无论是前进还是收缩后），我们都计算一下当前窗口的长度（right - left + 1），并和我们已知的 maxLen（最大长度）比较，始终保留那个较大的。

// 循环往复：

// 重复这个过程，直到 right 手指滑到字符串末尾。maxLen 就是最终答案。

func lengthOfLongestSubstring(s string) int {
	check := map[rune]int{}
	maxlen := 0
	left := 0
	for right, char := range s{
		if index,found := check[char];found && index >= left{
			left = index + 1 
		} 
		check[char] = right
		currentlen := right - left + 1
		if currentlen > maxlen {
			maxlen = currentlen
		}
	}
	return maxlen
}

func main(){
	s := "bbbbbbb"
	max_len := lengthOfLongestSubstring(s)
	fmt.Printf("the longestlen is : %d\n",max_len)
	s = "abc"
	max_len = lengthOfLongestSubstring(s)
	fmt.Printf("the longestlen is : %d",max_len)
}