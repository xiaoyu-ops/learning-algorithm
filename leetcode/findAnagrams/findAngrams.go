package main

import(
	"fmt";
	"reflect";
)
// 本题主要还是滑动窗口
// 您在代码中用两种完全不同的方式来“读取”字符串 s 中的字符，导致了类型不匹配。

// for right, var := range s

// 当您用 for range 遍历一个 string 时，Go 语言会智能地将其解码为 rune。

// rune 是 Go 用来表示一个 Unicode 字符的类型（它是一个 int32）。它可以代表 'a' (1个字节)，也可以代表 '你' (3个字节)。

// 所以，您的 var 变量是 rune 类型。您把它存入 check_s (类型为 map[rune]int)，这是正确的。

// check_s[s[left]]--

// 当您用 s[index]（比如 s[left]）来索引一个 string 时，Go 语言会执行字节索引。

// s[left] 返回的是在那个位置上的单个 byte（它是一个 uint8）。

// 这不是一个字符，只是一个字节。
func findAnagrams(s string, p string) []int {
	var ans []int
	check_p := make(map[rune]int)
	check_s := make(map[rune]int)
	sRunes := []rune(s)
	left := 0

	for _,y := range p{
		check_p[y]++
	}

	for right, v := range sRunes{
		check_s[v]++

		if right - left + 1 > len(p){
			leftChar := sRunes[left]
			check_s[leftChar]--
			if check_s[leftChar] == 0 {
				delete(check_s, leftChar) // 关键：移除数量为0的键
			}
			left++
		}
		
		if reflect.DeepEqual(check_s, check_p) {
			ans = append(ans,left)
		}
	}
	return ans
}	

func main() {
	s := "cbaebabacd"
	p := "abc"
	ans := findAnagrams(s,p)
	fmt.Printf("The answer is : %v",ans)
}