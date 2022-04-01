package main

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
)

// 自定义类型
// 相当于一个新类型
type defInt int

// 可以为这个自定义类型添加方法
func (i defInt) show()  {
	fmt.Println(i)
}

// 类型别名, myInt实际就是int类型, 编译后就是int
// 帮助代码编写更清晰
type myInt = int

// 结构类型
type Person struct {
	name string
	age int
	m map[int]int
}

type Stu struct {
	Sex byte `json:"sex" db:"sex"`
	Name int
	Person
}

func changeStruct(person Person)  {
	person.name = "修改"
}


func (p Person)funcRevicer() {
	p.age = 1
}

func (p* Person)ptrRevicer() {
	p.name = "ptr"
}

func main() {
	stu := Stu{}
	stu.age = 1
	stu.Name = 1
	stu.Person.name = "name"
	stu.ptrRevicer()
	fmt.Println(stu) // {0 1 {ptr 1 map[]}}

	jsonStr, _ := json.Marshal(stu)
	fmt.Printf("%#v\n",string(jsonStr))


	p1 := Person{}
	p2 := p1 // 会将 所有值拷贝一份 (深度的拷贝)
	p1.name = "123"
	p1.m = make(map[int]int, 2)
	p1.m[1] = 2
	changeStruct(p1)
	p1.funcRevicer()
	p1.ptrRevicer()
	fmt.Println(p1) // {123 0 map[1:2]}
	fmt.Println(p2) // { 0 map[]}
	fmt.Println(debug.Stack())


}

