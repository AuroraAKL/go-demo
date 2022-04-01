package main

import (
  "fmt"
  "strings"
  "time"
)

// 测试字符串拼接耗时
func testStringJoinSpendTime() string {
  start := time.Now().UnixNano() / 1e6
  str := "12312"
  for i:=0;i<20000;i++ {
    for j:=0; j< 5; j++ {
      str += fmt.Sprintf("%s %s %s", fmt.Sprintf("%d", 1123123), fmt.Sprintf("%s", "sddsdsfd"), fmt.Sprintf("%d", 123123123))
    }
  }
  fmt.Println(time.Now().UnixNano() / 1e6 - start)
  return str
}

// 测试字符串拼接耗时
func testStringBuilderSpendTime() string {
  start := time.Now().UnixNano() / 1e6
  var builder strings.Builder
  builder.WriteString("12312")
  for i:=0;i<20000;i++ {
    for j:=0; j< 5; j++ {
      s := fmt.Sprintf("%s %s %s", fmt.Sprintf("%d", 1123123), fmt.Sprintf("%s", "sddsdsfd"), fmt.Sprintf("%d", 123123123))
      s2 := fmt.Sprintf("%s %s %s", fmt.Sprintf("%d", 1123123), fmt.Sprintf("%s", "sddsdsfd"), fmt.Sprintf("%d", 123123123))
      builder.WriteString(s)
      builder.WriteString(s2)
    }
  }
  fmt.Println(time.Now().UnixNano() / 1e6 - start)
  return builder.String()
}

func TestStringLen() {
  str := "你好世界"
  rStr := []rune(str)
  fmt.Println(len(str), len(rStr))
}

func main() {
  TestStringLen()
  fmt.Println("0000" > "1110", "1" > "010")

  testStringBuilderSpendTime()
  testStringJoinSpendTime()



}
