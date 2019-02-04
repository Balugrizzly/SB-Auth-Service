# SB-Auth-Service
Authentication for SpeedBlocks

### General API Information
- All endpoints accept a JSON User struct/object.
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

### Issues
- Type conversion/casting in the decoding function is not working if only numbers are provided for string fields it will result in an empty string
