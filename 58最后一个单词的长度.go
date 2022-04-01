package main

import "fmt"

// 给你一个字符串 s，由若干单词组成，单词之间用空格隔开。返回字符串中最后一个单词的长度。如果不存在最后一个单词，请返回 0 。
//
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/length-of-last-word
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func lengthOfLastWord(s string) int {
	len := len(s)
	count := 0
	for i := len - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return count
		}
		count++
	}
	return count
}

func main() {
	i := lengthOfLastWord("Hello world")
	fmt.Println(i)

	i2 := lengthOfLastWord(" ")
	fmt.Println(i2)

	i3 := lengthOfLastWord("a ")
	fmt.Println(i3)
}