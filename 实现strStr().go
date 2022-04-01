package main

import "fmt"

// 实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-strstr
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func strStr(haystack string, needle string) int {
	nLen := len(needle)
	hLen := len(haystack)
	if nLen == 0 && hLen == nLen {
		return 0
	}
	for i := 0; i <= hLen - nLen; i++  {
		count := 0
		for j := 0; j < nLen; j++ {
			if haystack[i + count] == needle[j] {
				count++
			}
		}
		if count == nLen {
			return i
		}
	}
	return -1
}

func main() {
	i := strStr("123", "123")
	fmt.Println(i)
}