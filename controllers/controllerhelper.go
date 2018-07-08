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

func genericQueryGet(
  w http.ResponseWriter, r *http.Request, model interface{},
  allowedSearchKey []string) {
  w.Header().Set("Content-Type", "application/json")

  var queryValueMap = r.URL.Query()
  var queryMap = map[string]interface{}{}

  var isValid = false
  if len(queryValueMap) > 0 {
    for k, v := range queryValueMap{
      isValid = false
      for _, key := range allowedSearchKey {
        if k == key {
          queryMap[k] = v[0]
          isValid = true
          break
        }
      }
      if !isValid {
        break
      }
    }
  }

  if !isValid {
    var errorMessage = "Key is unknown or not allowed"
    var responseMap = map[string]string {
      "message": errorMessage,
    }
    w.WriteHeader(http.StatusNotAcceptable)
    json.NewEncoder(w).Encode(responseMap)
    return
  }

  var connectionInfo = db.Where(queryMap).Find(&model)

  if connectionInfo.Error != nil {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(constEMPTY)
  } else {
    json.NewEncoder(w).Encode(model)
  }
}

func genericQueryUpdate(
  w http.ResponseWriter, r *http.Request, model interface{},
  primaryKeys []string) {
  w.Header().Set("Content-Type", "application/json")

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

  var modelStruct = structToMap(model)
  var queryMap = map[string]interface{}{}
  var queryUpdate = map[string]interface{}{}

  var primaryKeySatisfied = 0
  for key, value := range modelStruct {
    var inPrimaryKey = false
    for _, k := range primaryKeys {
      if key == k {
        inPrimaryKey = true
        break
      }
    }

    if inPrimaryKey {
      primaryKeySatisfied += 1
      queryMap[key] = value
    } else {
      queryUpdate[key] = value
    }
  }

  if len(primaryKeys) != primaryKeySatisfied {
    var errorMessage = "Not All Primary Key Satisfied"
    var responseMap = map[string]string {
      "message": errorMessage,
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(responseMap)
    return
  }

  var noRecordError = db.Where(queryMap).Find(model).Error
  if noRecordError != nil {
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(constEMPTY)
  }

  var connectionInfo = db.Model(&model).Updates(queryUpdate)

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

func structToMap(sourceStruct interface{}) map[string]interface{} {
  var resultMap map[string]interface{}
	marshaledStruct, _ := json.Marshal(sourceStruct)
	json.Unmarshal(marshaledStruct, &resultMap)

  return resultMap
}
