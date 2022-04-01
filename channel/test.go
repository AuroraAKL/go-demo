package main

import (
  "fmt"
  "runtime"
  "sync"
)



func producer(c chan int, n int, w *sync.WaitGroup)  {
  for n > 0 {
    fmt.Println("producer", n)
    go func(n int) {
      c <- n
    }(n)
    n--
  }
  c <- -999
  w.Done()
}

func consumer(c chan int, w *sync.WaitGroup)  {
  fmt.Println("comsumer start")
  for {
    select {
    case i := <- c:
      if i == -999 {
        w.Done()
        fmt.Println("comsumer end")
        return
      }
      go fmt.Println(i)
    }
  }
}

func main() {
  w := sync.WaitGroup{}
  defer func() {
    if p := recover(); p != nil {
      fmt.Println(p)
    }
  }()
  fmt.Println("main start")
  c := make(chan int, 10)
  w.Add(2)
  go producer(c, 20, &w)
  go consumer(c, &w)
  //for runtime.NumGoroutine() > 2 {
  //  fmt.Println("NumGoroutine", runtime.NumGoroutine())
  //}
  w.Wait()
  fmt.Println("main end", runtime.NumGoroutine())
}


