package main

import (
  "fmt"
  "log"
  "net/http"
//  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "pmapi/versionOne"
)

func authenticationMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    next.ServeHTTP(w, r)
  })
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}


func main(){
  router := versionOne.NewRouter()

  allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
  allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
  // Launch server with CORS validations
  log.Fatal(http.ListenAndServe(":8008",handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
