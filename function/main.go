package main

import "fmt"

// 命名返回值, 可以省略return的变量
func returnFunc() (ret int) {
	ret = 10
	return
}

// 多个参数类型相同可以省略
func returnFunc2(x, y, z int, m, n string) (ret int) {
	ret = 10
	return
}

// 变长参数类型 必须为最后一个参数, arr类型是切片
func returnFunc3(x, y, z int, m, n string, arr... int) (ret int) {
	ret = 10
	return
}

func returnFunc4(arr... map[int]string) (ret int) {
	ret = len(arr)
	return
}

func deferFunc() [2]int {
	arr := [...]int{1, 2}
	defer func() {
		arr[0] = 0
	}()
	return arr
}
//
func deferFunc2() (arr [2]int) {
	arr = [...]int{1, 2}
	defer func() {
		arr[0] = 0
	}()
	return arr
}

// 函数 闭包
func pack(x int) (func() int, func() int)  {
	return func() int {
		x--
		return x
	},
	func() int {
			x++
			return x
		}
}


func main() {
	defer func() {
		fmt.Println("defer main")
	}()
	fmt.Println(deferFunc())
	fmt.Println(deferFunc2())
	fmt.Println("结束")

	// 函数闭包 = 函数 + 对外部变量的引用
	dec, inc := pack(100)
	fmt.Println(dec())
	fmt.Println(inc())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
}
