package versionOne

import (
  "fmt"
  "log"
  "io"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

type Controller struct {
    Repository Repository
}

type Response struct {
	Message string `json:"message,omitempty"`
	Object interface{} `json:"object,omitempty"`
}

func ResponseData(r Response) []byte {
  result, err := json.Marshal(r)
  if err != nil {
    return nil
  }
  return result
}


func (c *Controller) Hello(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello")

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(http.StatusOK)

  var currentSession = c.Repository.CreateSesion()
  response := Response{"Create session OK",currentSession}
  result := ResponseData(response)
  //result, err := json.Marshal(response)
  //if err != nil {
  //  http.Error(w, err.Error(), http.StatusInternalServerError)
  //  return
  //}
  w.Write(result)
  return
}


func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
  log.Println("CreateUser")

  var user User
  message := "Create user success"
  headerStatus := http.StatusCreated // default

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
  if err != nil {
    log.Fatalln("Error CreateUser", err)
    headerStatus = http.StatusInternalServerError
  }
  if err := r.Body.Close(); err != nil {
    log.Fatalln("Error CreateUser", err)
  }
  if err := json.Unmarshal(body, &user); err != nil { // unmarshall body contents as a type Candidate
    log.Println(err)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Fatalln("Error CreateUser unmarshalling data", err)
      headerStatus = http.StatusInternalServerError
    }
  }
//  log.Println(user)
  createdUser,err := c.Repository.InsertUser(user) // adds the user to the DB
  if err != nil {
    headerStatus = http.StatusInternalServerError
    message = "ERROR: Create user failed"
    log.Fatalln("Error CreateUser unmarshalling data", err)
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(headerStatus)

  response := Response{message, createdUser}
  data, _ := json.Marshal(response)
  w.Write(data)

  return
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
  var user User
//  headerStatus := http.StatusOK
  message := "Login successful"
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  sessionId := r.Header.Get("SessionID")

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
  if err != nil {
    log.Fatalln("Error Login", err)
  }
  if err := r.Body.Close(); err != nil {
    log.Fatalln("Error Login", err)
  }
  if err := json.Unmarshal(body, &user); err != nil { // unmarshall body contents as a type Candidate
    log.Println(err)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Fatalln("Error Login unmarshalling data", err)
    }
  }
  loginedUser,err := c.Repository.Login(user, sessionId)
  if err != nil {
    // error handle
    message = fmt.Sprintf("%v", err)
    //
    w.WriteHeader(http.StatusInternalServerError)
    data, _ := json.Marshal(Response{message, nil})
    w.Write(data)
    return
  }

  //UpdateSession(db, sessionId, loginedUser.ID)
  w.WriteHeader(http.StatusOK)
  response := Response{message, loginedUser}
  data, _ := json.Marshal(response)
  w.Write(data)

  return
}
