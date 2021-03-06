package main

import "fmt"

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。

// searchInsert
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v ==  target {
			return i
		}
		if v > target {
			return i
		}
	}
	return len(nums)
}

func main() {
	i := searchInsert([]int{1,3,5,6}, 8)
	fmt.Println(i)

}