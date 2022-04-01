package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type indexHandler2 struct {
  content string
}

func (ih *indexHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request )  {
  fmt.Fprintln(w, ih.content)
  ok, _ := json.Marshal(r.Header)
  fmt.Fprintln(w, string(ok))
  fmt.Fprintln(w, r.Header.Get("Cookie"))
  fmt.Fprintln(w, r.RequestURI, r.URL, r.Method)
}

func main() {
  http.Handle("/", &indexHandler2{content: "hello world!"})
  err := http.ListenAndServe(":8010", nil)
  if err != nil {
    return 
  }
}