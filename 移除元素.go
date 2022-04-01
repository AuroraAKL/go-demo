package main

import "fmt"

// 给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// swap 交换两个整数值
func swap(a* int, b*int){
	temp := *a
	*a = *b
	*b = temp
}
func removeElement(nums []int, val int) int {
	// 思路, 区间划分
	// 将 != val的元素划分到前面区域S1, == val的元素划分到后面区域S2
	// 使用index指向分界处, 指向S1的最后一个元素, 即S2的第一个元素前.
	// 长度=index+1
	lenIndex := -1
	for i, v := range nums {
		if v != val{
			lenIndex++
			// 交换
			swap(&nums[i], &nums[lenIndex])
		}
	}
	return lenIndex + 1
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 2, 3}
	swap(&arr[1], &arr[2])
	fmt.Println(arr)

}