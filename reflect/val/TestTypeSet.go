package main

import (
  "fmt"
  "reflect"
)

func fun(t... interface{})  {
  n := len(t)

  for i := 0; i < n; i++ {
    s := string(i)
    val2 := reflect.ValueOf(s)
    val := reflect.ValueOf(t[i])
    if !val.CanSet() {
      val = reflect.ValueOf(&t[i])
    }
    val = reflect.Indirect(val)
    val.Set(val2)

    fmt.Printf("%v \n", val.Interface())
  }
}

func main() {
  //t := make([]interface{}, 0)
  //t = append(t, "123", nil)
  //fun(t...)
  //fmt.Println(t)
  //
  //t2 := make([]interface{}, 0)
  //s := "123"
  //t2 = append(t2, &s, nil, 23423)
  //fun(t2...)
  //fmt.Println(t2)
  //
  //s2 := "123"
  //fun(s2)
  //fmt.Println(s2)

  s3 := "123"
  fun(&s3)
  fmt.Println(s3)

  s4 := "123"
  fun(s4)
  fmt.Println(s4)
}
