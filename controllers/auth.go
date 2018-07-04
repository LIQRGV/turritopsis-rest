package controllers

import (
  "os"
  "io"
  "net/http"
  "turritopsis-rest/models"
  "encoding/json"
  "golang.org/x/crypto/bcrypt"
  "github.com/gorilla/sessions"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "fmt"
)


func Login(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var secretKey = os.Getenv("SECRET_KEY")
  var key = []byte(secretKey)
  var store = sessions.NewCookieStore(key)
  var session, _ = store.Get(r, "go-cookie")
  var userId = session.Values["userId"]

  if userId != nil {
    // if user login already
    // might redirected later, but i return for now
    return
  }

  type acceptedSchema struct {
    Username string
    Password string
  }

  var jsonBuff acceptedSchema

  var isInputAccepted = validateInput(r.Body, &jsonBuff)

  if !isInputAccepted {
    http.Error(w, "{}", 401)
    return
  }

  var isUserValidated, user = validateUser(jsonBuff.Username, jsonBuff.Password)

  if !isUserValidated {
    http.Error(w, "{}", 401)
    return
  }

  session.Values["userId"] = user.Id
  session.Save(r,w)
  w.Write([]byte("{}"))
}

func validateInput(body io.ReadCloser, acceptedSchema interface{}) (bool) {
  var decoder = json.NewDecoder(body)
  var decodeError = decoder.Decode(acceptedSchema)

  return decodeError != nil
}

func validateUser(username string, password string) (bool, models.User) {
  var user models.User

  var dbHost = os.Getenv("DATABASE_HOST")
  var dbDriver = os.Getenv("DATABASE_DRIVER")
  var dbName = os.Getenv("DATABASE_NAME")
  var dbUsername = os.Getenv("DATABASE_USERNAME")
  var dbPassword = os.Getenv("DATABASE_PASSWORD")

  var dbUri = fmt.Sprintf(
    "%s://%s:%s@%s/%s", dbDriver, dbUsername, dbPassword, dbHost, dbName,
  )

  var db, _ = gorm.Open(
    dbDriver,
    dbUri,
  )

  db.Where("username = ?", username).Find(&user)

  return bcrypt.CompareHashAndPassword(
    []byte(user.Password), []byte(password),
  ) == nil, user
}
