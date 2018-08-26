package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/liqrgv/turritopsis-rest/models"
)

func ShowUsers(w http.ResponseWriter, r *http.Request) {
  var models []models.User

  db.Find(&models)

  json.NewEncoder(w).Encode(models)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  var model models.User
  var primaryKey = "id"

  genericGet(w, r, &model, primaryKey)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
  var model models.User
  var primaryKey = "id"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  var model models.User
  var primaryKey = "id"

  genericUpdate(w, r, &model, primaryKey)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  var model models.User
  var primaryKey = "id"

  genericDelete(w, r, &model, primaryKey)
}
