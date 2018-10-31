package versionOne

import (
  "errors"
  "fmt"
  "time"
  "os"
  "log"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/google/uuid"
)

type Repository struct{}

type Response struct {
  Message string `json:"message"`
}

const DB_NAME string = "./blog.v10.db"

func RemoveDatabase() {
  os.Remove(DB_NAME)
}

func InitialDatabase() {
  db, err := sql.Open("sqlite3", DB_NAME)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  sqlStmt := `
  CREATE TABLE IF NOT EXISTS User (ID text not null primary key, username text, password text, enabled bool, email text, introdution text);
  CREATE TABLE IF NOT EXISTS Session (ID text not null primary key, Owner text, CreatedAt integer);
  `
  _, err = db.Exec(sqlStmt)
  if err != nil {
    log.Printf("%q: %s\n", err, sqlStmt)
    return
  }
  return
}

func (r Repository) CreateSesion() Session {
  db, err := sql.Open("sqlite3", DB_NAME)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  tx, err := db.Begin()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := tx.Prepare("insert into Session(ID, Owner, CreatedAt) values(?, ?, ?)")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()
	//
  var createAt = time.Now().Unix()
  uuidObject,err := uuid.NewRandom()
  if err != nil{
      fmt.Println("Cannot create sessionId")
  }

  _, err = stmt.Exec(uuidObject.String(), "admin" ,createAt)
  if err != nil {
    log.Fatal(err)
  }

  tx.Commit()

  stmt, err = db.Prepare("select ID, Owner, CreatedAt from Session where ID = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()
  var sessionId string
  var owner string
  var createdat int64

  err = stmt.QueryRow(uuidObject.String()).Scan(&sessionId, &owner, &createdat) //
  if err != nil {
    log.Fatal(err)
  }

  session := Session{sessionId, owner, time.Unix(createdat, 0)}
  return session
}


func (r Repository) InsertUser(u User) (User, error){
  var user User
  // validate
  log.Println(u)
  if u.Username == "" ||u.Password == "" || u.Email == "" {
    return user, errors.New("User infomation is invalid")
  }
  if u.Introdution == "" {
    u.Introdution = " "
  }
  if u.Enabled == nil {
    u.Enabled = true
  }
  // connect to db
  db, err := sql.Open("sqlite3", DB_NAME)
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot insert user")
  }
  defer db.Close()

  tx, err := db.Begin()
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot insert user")
  }
  stmt, err := tx.Prepare("insert into User(ID, username, password, email, enabled, introdution) values(?, ?, ?, ?, ?, ?)")
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot insert user")
  }
  defer stmt.Close()

  uuidObject,err := uuid.NewRandom()
  if err != nil{
      fmt.Println("Cannot create user id")
      return user, errors.New("Cannot insert user")
  }

  _, err = stmt.Exec(uuidObject.String(), u.Username , u.Password, u.Email, u.Enabled, u.Introdution)
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot insert user")
  }

  tx.Commit()
  /**/
  stmt, err = db.Prepare("select id, username, email from User where ID = ? ")
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot get inserted user infomation")
  }
  defer stmt.Close()

  err = stmt.QueryRow(uuidObject.String()).Scan(&user.ID, &user.Username, &user.Email, &user.Introdution) //
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot get inserted user infomation")
  }

  /**/

  return user, nil
}

func (r Repository) Login(u User) bool {
  db, err := sql.Open("sqlite3", DB_NAME)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  //
  stmt, err := db.Prepare("select ID from User where username = ? and password = ? ")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  var userId string

  err = stmt.QueryRow(u.Username, u.Password).Scan(&userId) //
  if err != nil {
    log.Fatal(err)
  }
  //
  return true
}
