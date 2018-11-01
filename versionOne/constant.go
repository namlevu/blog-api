package versionOne

const DB_NAME string = "./blog.v10.db"
const SQLITE = "sqlite3"
const CONNECT_DB_ERROR_MSG = "Cannot connect with DB"


const SELECT_USER_BY_ID = "select * from User where ID = ? "
const SELECT_USER_BY_NAME = "select * from User where username = ? "

const SELECT_USER_FAILED = "Cannot get user infomation"
const INSERT_USER_STMT = `insert into User
													(ID, username, password, email, enabled, introdution)
													values(?, ?, ?, ?, ?, ?)`
const INSERT_USER_FAILED_MSG = "Cannot insert user"
const LOGIN_FAILED_MSG = "Cannot login"
const CREATE_DATABASE_STMT =  `
														CREATE TABLE IF NOT EXISTS User(
															ID text not null primary key,
															username text,
															password text,
															enabled bool,
															email text,
															introdution text
														);
														CREATE TABLE IF NOT EXISTS Session (
															ID text not null primary key,
															Owner text,
															CreatedAt integer
														);
														`
const INSERT_SESSION_STMT = "insert into Session(ID, CreatedAt) values(?, ?)"
const INSERT_SESSION_FAILED_MSG = "Cannot create session"
const USER_INFO_INVALID_MSG = "User infomation is invalid"
