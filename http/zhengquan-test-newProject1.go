package main

import (
"fmt"
"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world")
}

func main() {
  http.HandleFunc("/", indexHandler)
  err := http.ListenAndServe(":8009", nil)
  if err != nil {
    return 
  }
}