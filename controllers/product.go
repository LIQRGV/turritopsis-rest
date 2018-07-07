package controllers

import (
  "net/http"
  "encoding/json"
  "turritopsis-rest/models"
)

func ShowProducts(w http.ResponseWriter, r *http.Request) {
  var products []models.Product

  db.Find(&products)

  json.NewEncoder(w).Encode(products)
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
