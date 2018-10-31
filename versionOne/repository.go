package versionOne

import (
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

func InitialDatabase() {
  os.Remove(DB_NAME)

  db, err := sql.Open("sqlite3", DB_NAME)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  sqlStmt := `
  create table User (ID text not null primary key, username text, password text, enabled bool, email text, introdution text);
  delete from User;
  create table Session (ID text not null primary key, Owner text, CreatedAt integer);
  delete from Session;
  `
  _, err = db.Exec(sqlStmt)
  if err != nil {
    log.Printf("%q: %s\n", err, sqlStmt)
    return
  }
  return
}

func CreateSesion() Session {
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

func InsertUser() User{
  var user User
  // TODO:
  return user
}
