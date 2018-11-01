package versionOne

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type Repository struct{}

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
	db, err := sql.Open(SQLITE, DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(CREATE_DATABASE_STMT)
	if err != nil {
		log.Printf("%q: %s\n", err, CREATE_DATABASE_STMT)
		return
	}
	return
}

func IsUserValid(u User) bool {
	if u.Username == "" || u.Password == "" || u.Email == "" {
		return false
	}
	return true
}

func (r Repository) CreateSesion() string {
	db, err := sql.Open(SQLITE, DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(INSERT_SESSION_STMT)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var createAt = time.Now().Unix()
	sessionIdObj, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(INSERT_SESSION_FAILED_MSG)
	}

	_, err = stmt.Exec(sessionIdObj.String(), createAt)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	return sessionIdObj.String() // sessionId
}

func GetUserById(db DB, userId string) (User, error) {
	var user User
	stmt, err := db.Prepare(SELECT_USER_BY_ID)
	if err != nil {
		log.Fatal(err)
		return user, errors.New(SELECT_USER_FAILED)
	}
	defer stmt.Close()

	err = stmt.QueryRow(userId).Scan(&user.ID,
																&user.Username,
																&user.Email,
																&user.Enabled,
																&user.Introdution)
	if err != nil {
		log.Fatal(err)
		return user, errors.New(SELECT_USER_FAILED)
	}

	return user, nil
}

func (r Repository) InsertUser(u User) (User, error) {
	var user User

	if !IsUserValid(u) {
		return user, errors.New(USER_INFO_INVALID_MSG)
	}

	db, err := sql.Open(SQLITE, DB_NAME)
	if err != nil {
		log.Fatal(err)
		return user, errors.New(INSERT_USER_FAILED_MSG)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return user, errors.New(INSERT_USER_FAILED_MSG)
	}
	stmt, err := tx.Prepare(INSERT_USER_STMT)
	if err != nil {
		log.Fatal(err)
		return user, errors.New(INSERT_USER_FAILED_MSG)
	}
	defer stmt.Close()

	uuidObject, err := uuid.NewRandom()
	if err != nil {
		return user, errors.New(INSERT_USER_FAILED_MSG)
	}
	hashpassword, _ := HashPassword(u.Password)

	_, err = stmt.Exec(uuidObject.String(),
											u.Username,
											hashpassword,
											u.Email,
											u.Enabled,
											u.Introdution)
	if err != nil {
		log.Fatal(err)
		return user, errors.New(INSERT_USER_FAILED_MSG)
	}

	tx.Commit()

	user, err = GetUserById(db, uuidObject.String())

	return user, nil
}

func (r Repository) Login(u User) (User, error) {
	var user User

	db, err := sql.Open(SQLITE, DB_NAME)
	if err != nil {
		log.Fatal(err)
		return User{}, errors.New(CONNECT_DB_ERROR_MSG)
	}
	defer db.Close()
	//
	stmt, err := db.Prepare(SELECT_USER_BY_NAME)
	if err != nil {
		log.Fatal(err)
		return User{}, errors.New(SELECT_USER_FAILED)
	}
	defer stmt.Close()

	err = stmt.QueryRow(u.Username).Scan(&user.ID,
																				&user.Username,
																				&user.Password,
																				&user.Enabled,
																				&user.Email,
																				&user.Introdution) //
	if err != nil {
		log.Fatal(err)
		return User{}, errors.New(SELECT_USER_FAILED)
	}

	if CheckPasswordHash(u.Password, user.Password) {
		if user.Enabled {
			return user, nil
		}
	}
	return User{}, errors.New(LOGIN_FAILED_MSG)
}
