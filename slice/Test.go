package main

import "fmt"

func main() {

  s := []int64{1, 2, 3, 4}

  s = append(s[:0], s[1:]...)

  fmt.Println(s)

}
