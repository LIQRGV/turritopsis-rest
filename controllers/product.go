package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/liqrgv/turritopsis-rest/models"
)

func ShowProducts(w http.ResponseWriter, r *http.Request) {
  var models []models.Product

  db.Find(&models)

  json.NewEncoder(w).Encode(models)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
  var model models.Product
  var primaryKey = "code"

  genericGet(w, r, &model, primaryKey)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
  var model models.Product
  var primaryKey = "code"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
  var model models.Product
  var primaryKey = "code"

  genericUpdate(w, r, &model, primaryKey)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
  var model models.Product
  var primaryKey = "code"

  genericDelete(w, r, &model, primaryKey)
}
