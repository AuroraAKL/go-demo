package main

import (
  "fmt"
  "os/exec"
  "sync"
  "time"
)

func testCmd() {
  cmd := exec.Command("notepad")
  cmd.Start()
  time.Sleep(3000)
  //fmt.Println(pid, cmd.ProcessState.Exited())
  cmd.Process.Kill()
  //fmt.Println("Exit Code", cmd.ProcessState.ExitCode())
}

func testGoroutine() {
  cmd := exec.Command("notepad")
  wait := sync.WaitGroup{}
  wait2 := sync.WaitGroup{}
  wait.Add(1)
  wait2.Add(1)
  go func() {
    cmd.Start()
    fmt.Println("start success")
    wait.Done()
    cmd.Wait()
    fmt.Println("start end")
    wait2.Done()
  }()

  wait.Wait()
  time.Sleep(30 * time.Second)
  cmd.Process.Kill()
  wait2.Wait()
  time.Sleep(60 * time.Second)
  fmt.Println("testGoroutine end")
}

func main() {
  testGoroutine()
}
