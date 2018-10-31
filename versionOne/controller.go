package versionOne

import (
  "log"
  "io"
  "io/ioutil"
  "net/http"
  "encoding/json"
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

  var currentSession = c.Repository.CreateSesion()
  helloResp := HelloResponse{"Create session OK",currentSession}

  result, err := json.Marshal(helloResp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(result)
  return
}

type UserResponse struct {
  Message string
  User User
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
  log.Println("CreateUser")
  var user User
  /**/
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
  if err != nil {
    log.Fatalln("Error CreateUser", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  if err := r.Body.Close(); err != nil {
    log.Fatalln("Error CreateUser", err)
  }
  if err := json.Unmarshal(body, &user); err != nil { // unmarshall body contents as a type Candidate
    w.WriteHeader(422) // unprocessable entity
    log.Println(err)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Fatalln("Error CreateUser unmarshalling data", err)
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
  }

  log.Println(user)
  createdUser,err := c.Repository.InsertUser(user) // adds the user to the DB
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := UserResponse{Message: "ERROR: Create user failed"}
    data, _ := json.Marshal(response)
    w.Write(data)
    return
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  response := UserResponse{"Create user success", createdUser}
  data, _ := json.Marshal(response)
  w.Write(data)
  /**/
  return
}
