

| #  | API       | Method | Inputs       | Outputs     | Note                           |
| -- | --------- | ------ | ------------ | ----------- | ------------------------------ |
| 01 | /hello    | POST   | client type,<br>version | sessionId,<br>message  | client type: app/web. version of app of web client |
| 02 | /login    | POST   | username,<br>password,<br>sessionId | message   | login|
| 03 | /logout   | POST   | sessionId | message | logout |
| 04 | /users    | GET    | sessionId,<br>filter   | User list,<br>message  | get all users. (pagination) |
| 05 |           | POST   | sessionId,<br>User object | User obj,<br>message | add new user |
| 06 |           | PUT    | sessionId,<br>User object | User obj,<br>message | update user info |
| 07 |           | DELETE | sessionId,<br>User object | message  | update user info |
| 08 | /users/id | GET    | sessionId,<br>User id | User obj,<br>message | get user info by id            |
| 09 | /blogs    | GET    | sessionId,<br>filter | blog list,<br>message | filter ( author, date, category) |
| 10 |           | POST   | sessionId,<br>Blog obj | Blog obj,<br>message ||
| 11 |           | PUT    | sessionId,<br>Blog obj | Blog obj,<br>message ||
| 12 |           | DELETE | sessionId,<br>Blog obj    | message | update user info |
| 13 | /blogs/   | GET    | sessionId,<br>slug,<br>blog id | Blog obj,<br>message  | /blogs/slug/id |
| 14 | /bye      | POST   | sessionId | message ||
