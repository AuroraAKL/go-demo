package main

import (
	"fmt"
	"unicode"
)

func testInterface() {
}

func testStruct() {
	type mytype struct {
		name     int
		password string
	}
	t := mytype{password: "sdfsd"}
	t.name = 123
	fmt.Printf("%s %d", t.password, t.name)
	fmt.Printf("%v", t)
}

func testArray() {
	arr := [5]string{}
	arr[1] = `123123`
	for _, item := range arr {
		fmt.Println(item)
	}
}

func testSlice() {
	arr := make([]int, 3)
	arr[0] = 123
	arr[2] = 12311
	arr = append(arr, 999)
	arr = append(arr, 90909, 123, 12323, 123)
	for index, item := range arr[:3] {
		fmt.Printf("%d %d\n", index, item)
	}
	fmt.Printf("len=%d cap=%d", len(arr), cap(arr))
}

type MyType struct {
	name string
}

// 测试指针变量
func testPtr(ptr *MyType) {
	str := MyType{name: "sfdsaf"}
	*ptr = str
}

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


func main() {
	var arr = [5]string{}
	arr2 := arr
	fmt.Println(arr)
	fmt.Println(arr2)

	if unicode.Is( unicode.Han, '杭'){
		fmt.Println("汉子")
	}

	//fmt.Printf("test\n")

	//testSlice()

	//testArray()

	//testStruct()

	myType := MyType{name: "原"}
	p := &myType
	testPtr(p)
	fmt.Println(myType.name)
	fmt.Println(p.name)
	//
	//fmt.Println(len([]rune("123哈哈"))) // 5
	//fmt.Println(len("123哈哈"))         //  9
	//fmt.Printf("%T\n", 'h')           // int32


	//arr := [6] int64{1, 2, 3, 4, 5, 6}
	//slice := arr[:]
	//arr[0] = -1
	//fmt.Println(slice) // [-1 2 3 4 5 6]
	//
	//s1 := make([]int, 5, 10) // size:5, cap: 10
	//for i, v := range s1 {
	//	fmt.Println(i, v)
	//}
	//
	//// 将 下标为 2的元素删除,  ...表示将数组拆开
	//s1 = append(s1[:2], s1[3:]...)
	//
	//sort.Ints(s1)
	//fmt.Println(s1)
	//
	//var a = new(int)

}
