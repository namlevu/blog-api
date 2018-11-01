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
  "golang.org/x/crypto/bcrypt"
)

type Repository struct{}

//type Response struct {
//  Message string `json:"message"`
//}

const DB_NAME string = "./blog.v10.db"

func RemoveDatabase() {
  os.Remove(DB_NAME)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
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
  log.Println("repository", u)
  if u.Username == "" ||u.Password == "" || u.Email == "" {
    return user, errors.New("User infomation is invalid")
  }
  if u.Introdution == "" {
    u.Introdution = " "
  }

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
  log.Println("Repository u.Enabled: ", u.Enabled)
  hashpassword,_ := HashPassword(u.Password)

  _, err = stmt.Exec(uuidObject.String(), u.Username, hashpassword, u.Email, u.Enabled, u.Introdution)
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot insert user")
  }

  tx.Commit()

  stmt, err = db.Prepare("select id, username, email, enabled, introdution from User where ID = ? ")
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot get inserted user infomation")
  }
  defer stmt.Close()

  err = stmt.QueryRow(uuidObject.String()).Scan(&user.ID, &user.Username, &user.Email, &user.Enabled, &user.Introdution) //
  if err != nil {
    log.Fatal(err)
    return user, errors.New("Cannot get inserted user infomation")
  }

  return user, nil
}

func (r Repository) Login(u User) (User, error) {
  var user User

  db, err := sql.Open("sqlite3", DB_NAME)
  if err != nil {
    log.Fatal(err)
    return User{}, errors.New("Cannot connect with DB")
  }
  defer db.Close()
  //
  stmt, err := db.Prepare("select * from User where username = ? ")
  if err != nil {
    log.Fatal(err)
    return User{}, errors.New("Cannot find user")
  }
  defer stmt.Close()

  err = stmt.QueryRow(u.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Enabled, &user.Email, &user.Introdution) //
  if err != nil {
    log.Fatal(err)
    return User{}, errors.New("Cannot retrieve data")
  }

  if CheckPasswordHash( u.Password, user.Password) {
    if user.Enabled {
      return user, nil
    }
  }
  return User{}, errors.New("Cannot login")
}
