package versionOne

import (
  "net/http"
  "encoding/json"
  "github.com/google/uuid"
)

type Controller struct {
    Repository Repository
}


func (c *Controller) Hello(w http.ResponseWriter, r *http.Request) {

  uuidObject,err := uuid.NewRandom()
  if err != nil{
      fmt.Println("Cannot create sessionId")
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(http.StatusOK)
  type HelloResponse struct {
    Message string
    SessionId string
  }

  helloResp := HelloResponse{"Welcome to PM API",uuidObject.String()}

  result, err := json.Marshal(helloResp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(result)
  return
}
