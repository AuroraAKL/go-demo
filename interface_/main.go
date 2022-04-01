package main

import "fmt"

type A struct {

}

type B struct {

}

type Show interface {

}

func (a A) show()  {

}

func (b B) show()  {
	fmt.Println(b)
}

func main() {

}
