package main

import (
  "fmt"
  "reflect"
  "time"
)

// not ok
func reversal(arr...interface{})  {
  for i , end := 0, len(arr) - 1; i <= end / 2; i++ {
    j := end - i
    t := arr[i]
    arr[i] = arr[j]
    arr[j] = t
  }
}

// 获取 []*struct类型
type Str struct {
  Id int64
}

func (s Str) ReturnFunc() []*Str {
  arr := make([]*Str, 0)
  arr = append(arr, &Str{Id: 10})
  return arr
}

func testSlicePtr()  {

  method := reflect.ValueOf(Str{}).MethodByName("ReturnFunc")
  fmt.Println(method, method.IsValid())
  ansVals := method.Call(make([]reflect.Value, 0))

  for _, ansVal := range ansVals {
    fmt.Println(ansVal.Kind(), ansVal.Type(), ansVal, ansVal.Interface())
  }

  slicePtrVal := ansVals[0]
  GetReflectSliceFieldValue(slicePtrVal)
}

func GetReflectSliceFieldValue(value reflect.Value) []interface{} {
  size := value.Len()
  var rl []interface{}
  for i := 0; i < size; i++ {
    val := value.Index(i)
    fmt.Println(val, val.Type(), val.Kind(), val.Kind() == reflect.Ptr)
    elem := val.Elem()
    fmt.Println(elem, elem.Type(), elem.Kind(), elem.Type().String(), elem.Type().String() == "[]uint8")
    rl = append(rl, val)
  }
  return rl
}

func testTime() {
  now := time.Now()
  fmt.Println(now)

  value := reflect.ValueOf(now)
  fmt.Println(now, value, value.Kind(), value.Kind() == reflect.Struct)
  switch value.Interface().(type) {
  case time.Time:
    fmt.Println(value)

  }
}

func main()  {
  //arr := []int{1, 2, 3, 4, 5}
  //reversal(arr)
  //fmt.Println(arr)

  //testSlicePtr()
  testTime()


}
