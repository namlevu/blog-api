package main

import (
  "fmt"
  "log"
  "net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("Logging start")
    log.Println(r.URL.Path)
    f(w, r)
    log.Println("Logging end")
  }
}

func authentication(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println("Authentication start")
    f(w, r)
    log.Println("Authentication end")
  }
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}


func main(){
  log.Println("[LOG] Main package run")
  fmt.Println("[CONSOLE] Main package run")
  //
  http.HandleFunc("/foo", authentication(logging(foo)))
  http.HandleFunc("/bar", authentication(logging(bar)))

  //
  http.ListenAndServe(":8008",nil)
}
