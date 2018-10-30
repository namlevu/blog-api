package versionOne

import (
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

func initialDatabase() {
  os.Remove(DB_NAME)

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table User (ID text not null primary key, username text, password text, enabled bool, email text, introdution text);
	delete from User;
  create table Session (  ID text not null primary key, Owner text, CreatedAt text);
  delete from Session;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
  return
}

func createSesion() Session {
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
  var createAt = time.Now().String()
  uuidObject,err := uuid.NewRandom()
  if err != nil{
      fmt.Println("Cannot create sessionId")
  }

  _, err = stmt.Exec(uuidObject.String(), "admin" ,createAt))
  if err != nil {
    log.Fatal(err)
  }

	tx.Commit()

  stmt, err = db.Prepare("select * from Session where ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var session Session
	err = stmt.QueryRow(uuidObject.String()).Scan(&Session)
	if err != nil {
		log.Fatal(err)
	}
  return session
}

func insertUser() User{
  db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
