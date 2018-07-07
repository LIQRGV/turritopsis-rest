package controllers

import (
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "os"
  "net/http"
  "reflect"
  "regexp"
  "strings"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error
var constEMPTY = map[string]string{}

var getAllWordRegex = regexp.MustCompile("([A-Za-z]+)")

func InitController() {
  var dbHost = os.Getenv("DATABASE_HOST")
  var dbDriver = os.Getenv("DATABASE_DRIVER")
  var dbName = os.Getenv("DATABASE_NAME")
  var dbUsername = os.Getenv("DATABASE_USERNAME")
  var dbPassword = os.Getenv("DATABASE_PASSWORD")

  var dbUri = fmt.Sprintf(
    "%s://%s:%s@%s/%s", dbDriver, dbUsername, dbPassword, dbHost, dbName,
  )

  db, err = gorm.Open(
    dbDriver,
    dbUri,
  )
}

// ====== PRIVATE FUNC

// ====== GENERAL HELPER
func validateInput(body io.ReadCloser, acceptedSchema interface{}) (bool) {
  var decoder = json.NewDecoder(body)
  var decodeError = decoder.Decode(acceptedSchema)

  return decodeError == nil
}

func toCamelCase(str string) string {
  var capitalizedWord = getAllWordRegex.ReplaceAllStringFunc(str, func(word string) string {
    return strings.Title(word)
  })

  return strings.Replace(capitalizedWord, "_", "", -1)
}

// ====== CONTROLLER HELPER
func genericGet(
  w http.ResponseWriter, r *http.Request, model interface{},
  primaryKey string) {
  w.Header().Set("Content-Type", "application/json")

  var whereClause = fmt.Sprintf("%s = ?", primaryKey)
  var params = mux.Vars(r)
  var code = params[primaryKey]

  var err = db.Where(whereClause, code).Find(model).Error
  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(constEMPTY)
  } else {
    json.NewEncoder(w).Encode(model)
  }

}

func genericCreate(
  w http.ResponseWriter, r *http.Request, model interface{},
  primaryKey string) {
  w.Header().Set("Content-Type", "application/json")

  var params = mux.Vars(r)
  var code = params[primaryKey]

  var isInputAccepted = validateInput(r.Body, model)
  if !isInputAccepted {
    var errorMessage = "Wrong Format JSON"
    var responseMap = map[string]string {
      "message": errorMessage,
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(responseMap)
    return
  }

  setPrimaryKey(model, primaryKey, code)

  var connectionInfo = db.Create(model)

  if connectionInfo.Error != nil {
    var errorMessage = fmt.Sprintf("%s", connectionInfo.Error)
    var responseMap = map[string]string {
      "message": errorMessage,
    }
    w.WriteHeader(http.StatusConflict)
    json.NewEncoder(w).Encode(responseMap)
    return
  }

  json.NewEncoder(w).Encode(model)
}

func genericUpdate(
  w http.ResponseWriter, r *http.Request, model interface{},
  primaryKey string) {
  w.Header().Set("Content-Type", "application/json")

  var whereClause = fmt.Sprintf("%s = ?", primaryKey)
  var params = mux.Vars(r)
  var code = params[primaryKey]

  var noRecordError = db.Where(whereClause, code).Find(model).Error
  if noRecordError != nil {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(constEMPTY)
  }

  var isInputAccepted = validateInput(r.Body, model)
  if !isInputAccepted {
    var errorMessage = "Wrong Format JSON"
    var responseMap = map[string]string {
      "message": errorMessage,
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(responseMap)
    return
  }

  setPrimaryKey(model, primaryKey, code)

  var connectionInfo = db.Save(model)

  if connectionInfo.Error != nil {
    var errorMessage = fmt.Sprintf("%s", connectionInfo.Error)
    var responseMap = map[string]string {
      "message": errorMessage,
    }
    w.WriteHeader(http.StatusConflict)
    json.NewEncoder(w).Encode(responseMap)
    return
  }

  json.NewEncoder(w).Encode(model)

}

func genericDelete(
  w http.ResponseWriter, r *http.Request, model interface{},
  primaryKey string) {
  w.Header().Set("Content-Type", "application/json")

  var whereClause = fmt.Sprintf("%s = ?", primaryKey)
  var params = mux.Vars(r)
  var code = params[primaryKey]

  var noRecordError = db.Where(whereClause, code).Find(model).Error
  if noRecordError != nil {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(constEMPTY)
  } else {
    db.Delete(model)
    json.NewEncoder(w).Encode(model)
  }
}

// ====== HACKISH UTIL

func setPrimaryKey(model interface{}, primaryKey string, code string) {
  var reflectedModel = reflect.ValueOf(model)
  if reflectedModel.Kind() != reflect.Ptr || reflectedModel.Elem().Kind() != reflect.Struct {
    errors.New("v must be pointer to struct")
  }
  var dereferModel = reflectedModel.Elem()
  var camelCasePrimary = toCamelCase(primaryKey)
  var reflectedAttribute = dereferModel.FieldByName(camelCasePrimary)

  switch reflectedAttribute.Kind() {
    case reflect.String:
      reflectedAttribute.SetString(code)
  }
}

func structToMap(sourceStruct interface{}) map[string]string {
  var resultMap map[string]string
	marshaledStruct, _ := json.Marshal(sourceStruct)
	json.Unmarshal(marshaledStruct, &resultMap)

  return resultMap
}
