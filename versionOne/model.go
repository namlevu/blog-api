package versionOne

import "time"

type User struct {
  ID string
  Username string
  Password string
  Enabled bool
  Email string
  Introdution string
}
type Session struct {
  ID string
  OwnerId string
  CreatedAt time.Time
}

type Tag struct {
  ID string
  Text string
}

type Comment struct {
  ID string
  Author string
  AuthorEmail string
  Content string
  Displayed bool
  CreatedAt time.Time
}

type Blog struct {
  ID string
  Title string
  Slug string
  RawContent string
  HtmlContent string
  Author User
  CreatedAt time.Time
  Tags []Tag
  Comments []Comment
}
