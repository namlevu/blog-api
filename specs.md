

| #  | API       | Method | input       | output     | note                           |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 01 | /hello    | POST   | client type | sessionId  | client type: app/web           |
|    |           |        | version     | message    | version of app of web client   |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 02 | /login    | POST   | username    | login OK   | login                          |
|    |           |        | password    | or NOK     |                                |
|    |           |        | sessionId   |            |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 03 | /logout   | POST   | sessionId   | OK         | logout                         |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 04 | /users    | GET    | sessionId   | OK or NOK  | get all users. (pagination)    |
|    |           |        | filter      | User list  |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 05 |           | POST   | sessionId   | OK or NOK  | add new user                   |
|    |           |        | User object | User obj   |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 06 |           | PUT    | sessionId   | OK or NOK  | update user info               |
|    |           |        | User object | User obj   |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 07 |           | DELETE | sessionId   | OK or NOK  | update user info               |
|    |           |        | User object | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 08 | /users/id | GET    |             | OK or NOK  | get user info by id            |
|    |           |        | User id     | User obj   |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 09 | /blogs    | GET    | sessionId   | OK or NOK  | filter ( author, date, cate)   |
|    |           |        | filter      | blog list  |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 10 |           | POST   | sessionId   | OK or NOK  |                                |
|    |           |        | Blog obj    | Blog obj   |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 11 |           | PUT    | sessionId   | OK or NOK  |                                |
|    |           |        | Blog obj    | Blog obj   |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 12 |           | DELETE | sessionId   | OK or NOK  | update user info               |
|    |           |        | User object | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 13 | /blogs/   | GET    | sessionId   | OK or NOK  | /blogs/<slug>/<id>             |
|    |           |        | slug        | Blog obj   |                                |
|    |           |        | blog id     | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
| 14 | /bye      | POST   | sessionId   | OK or NOK  |                                |
|    |           |        |             | message    |                                |
| -- | --------- | ------ | ----------- | ---------- | ------------------------------ |
