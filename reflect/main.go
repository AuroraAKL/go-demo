package main

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("init")
}
// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
func getType(x interface{})  {
	v := reflect.TypeOf(x)
	fmt.Printf("%v: %v,%v\n", v, v.Name(), v.Kind())
}

//
func getMethods(x interface{})  {
	xType := reflect.ValueOf(x)
	for i := 0; i < xType.NumMethod(); i++ {
		//m := xType.Method(i)
		params := make([]reflect.Value, 0)
		m := reflect.ValueOf(x).Method(i)
		fmt.Println("method=", m)
		m.Call(params)
	}
}

//func getMethod(x interface{})  {
//	xType := reflect.TypeOf(x)
//	fmt.Println("method=", m)
//	params := make([]reflect.Value, 0)
//	vx := reflect.ValueOf(x).Method()
//	params = append(params, vx)
//	m.Func.Call(params)
//}

type Struct struct {
	name string
}

func (s Struct) Show() {
	fmt.Println("show=", s.name)
}

func main() {
	fmt.Println("main")
	p := new(Struct)
	s := new(Struct)
	s.name = "s"
	getType(1)
	getType(Struct{})
	getType(p)

	fmt.Println("getMethods -------- ")
	getMethods(Struct{name: "s1"})
	getMethods(Struct{name: "s2"})
}
