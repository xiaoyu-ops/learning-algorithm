package main

import (
	"fmt";
	"sort"
)

// 注意题目会给我一个字符串数组，我需要将字母异位词组合
// 在一起，返回一个二维字符串数组
// 思路：使用哈希表，key为排序后的字符串，value
// 为字符串数组
// 遍历输入字符串数组，将每个字符串排序后作为key，
// 将原字符串添加到对应的value数组中


// 破题点1在于： abc和cba排序后都是abc
// 所以需要一个排序函数
func groupAnagrams(strs []string) [][]string {
    mp := map[string][]string{}
	for _, str := range strs{
		s := []byte(str)
		// 排序
		sort.Slice(s,func(i,j int) bool {return s[i]<s[j]})
		key := string(s)
		mp[key] = append(mp[key], str)
	}
	result := [][]string{}
	for _, v := range mp{
		result = append(result, v)
	}
	return result
}

// 破题点二：直接对字母进行计数
func groupAnagrams_2(strs []string) [][]string {
    mp := map[[26]int][]string{}
	for _, str := range strs{
		cnt := [26]int{}
		// 注意这里的[26]int{}
		// 是一个数组，包含26个槽位每个槽位中包含的
		// 是int类型的数字。
		for _, ch := range str{
			cnt[ch - 'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	result := [][]string{}
	for _, v := range mp{
		result = append(result, v)
	}
	return result
}
func main(){
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	answer := groupAnagrams(strs)
	fmt.Printf("the answer is :%v\n", answer)
	answer2 := groupAnagrams_2(strs)
	fmt.Printf("the answer2 is :%v", answer2)
}