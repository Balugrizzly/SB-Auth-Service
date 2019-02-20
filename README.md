# SB-Auth-Service
Authentication for SpeedBlocks

### General API Information
- All endpoints accept a JSON User object/struct in the request body and a JWtoken in the request header.
- All endpoints return a JSON object.
- Any endpoint can return an ERROR; the error payload is based on the ErrorResponse struct in responses.go and looks as follows:
`{
  "Status": false,
  "Msg": "optional message"
}`


### go gets
- go get -u github.com/gorilla/mux
- go get -u github.com/jinzhu/gorm
- go get github.com/mattn/go-sqlite3
- go get golang.org/x/crypto/bcrypt
- go get github.com/dgrijalva/jwt-go



## Endpoints

##### CreateUser
`POST /createuser`

Creates a user with the desired Name and Pw.

**Headers:** None

**Request Body:**

Name | Type | Mandatory | Info
------------ | ------------ | ------------
Name | STRING | YES | Has to be unique.
Pw | STRING | NO | Can be empty.


**Response:**
```json
{
  "ID": 4,
  "CreatedAt": "2019-02-20T17:14:46.9831406+01:00",
  "UpdatedAt": "2019-02-20T17:14:46.9831406+01:00",
  "DeletedAt": null,
  "Name": "Testuser",
  "Pw": "$2a$10$CjgurMstwGqPsF3chyVlseC8Pmx4/ytH0lCCAOHXZhV8g6Yon25Q6",
  "IsSuperuser": false
}
```

##### Authuser
`POST /authuser`

**Headers:** None

**Request Body:**

Name | Type | Mandatory |
------------ | ------------
Name | STRING | YES
Pw | STRING | YES

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkIjoxNTUwNjgxNzMyfQ.FysuE5BTAY-xW7c47CorKU0NCoSxjGOpzu43QA2JF64"
}
```

##### Isauthenticated
`POST /isauthenticated`

**Headers:**

Name | Type | Mandatory |
------------ | ------------
token | STRING | YES

**Request Body:** None


**Response:**
```json
{
  "status": true
}
```

##### Updateuser
`POST /updateuser`

Updates the user associated with the provided token.

Non given or emtpy Request Body parameters will not overwrite existing data!

**Headers:**

Name | Type | Mandatory |
------------ | ------------
token | STRING | YES

**Request Body:**

Name | Type | Mandatory |
------------ | ------------
Name | STRING | NO
Pw | STRING | NO

**Response:**
```json
{
  "ID": 4,
  "CreatedAt": "2019-02-20T17:14:46.9831406+01:00",
  "UpdatedAt": "2019-02-20T18:03:51.79157323+01:00",
  "DeletedAt": null,
  "Name": "TestuserChanged",
  "Pw": "$2a$10$CjgurMstwGqPsF3chyVlseC8Pmx4/ytH0lCCAOHXZhV8g6Yon25Q6",
  "IsSuperuser": false
}
```

##### Deleteuser
`POST /deleteuser`

Soft deletes the user associated with the provided token.

**Headers:**

Name | Type | Mandatory |
------------ | ------------
token | STRING | YES

**Request Body:** None

**Response:**
```json
{
  "status": true
}
```

##### Getuser
`Any /getuser`

Gets the user by id or name.
The Password wont be returned!

**Headers:** None

**Request Body:**

Name | Type | Mandatory |
------------ | ------------
Name | STRING | NO
Id | INT | NO


**Response:**
```json
{
  "ID": 1,
  "CreatedAt": "2009-02-07T03:51:41.116494428+01:00",
  "UpdatedAt": "2019-22-20T17:55:32.91364136+01:00",
  "DeletedAt": null,
  "Name": "user",
  "Pw": "",
  "IsSuperuser": false
}
```

##### Usernameisavailable
`Any /usernameisavailable`

**Headers:** None

**Request Body:**

Name | Type | Mandatory |
------------ | ------------
Name | STRING | NO

**Response:**
```json
{
  "status": true
}
```
