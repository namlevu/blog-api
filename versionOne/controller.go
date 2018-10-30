package versionOne

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/google/uuid"
)

type Controller struct {
    Repository Repository
}


func (c *Controller) Hello(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello")

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(http.StatusOK)
  type HelloResponse struct {
    Message string
    CurrentSession Session
  }

  var currentSession = createSesion()
  helloResp := HelloResponse{"Create session OK",currentSession}

  result, err := json.Marshal(helloResp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(result)
  return
}
