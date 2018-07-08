package controllers

import (
  "net/http"
  "encoding/json"
  "turritopsis-rest/models"
)

func ShowOrders(w http.ResponseWriter, r *http.Request) {
  var models []models.Order

  db.Find(&models)

  json.NewEncoder(w).Encode(models)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
  var model models.Order
  var primaryKey = "id"

  genericGet(w, r, &model, primaryKey)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
  var model models.Order
  var primaryKey = "id"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
  var model models.Order
  var primaryKey = "id"

  genericUpdate(w, r, &model, primaryKey)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
  var model models.Order
  var primaryKey = "id"

  genericDelete(w, r, &model, primaryKey)
}
