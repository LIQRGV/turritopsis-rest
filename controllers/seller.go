package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/liqrgv/turritopsis-rest/models"
)

func ShowSellers(w http.ResponseWriter, r *http.Request) {
  var models []models.Seller

  db.Find(&models)

  json.NewEncoder(w).Encode(models)
}

func GetSeller(w http.ResponseWriter, r *http.Request) {
  var model models.Seller
  var primaryKey = "id"

  genericGet(w, r, &model, primaryKey)
}

func CreateSeller(w http.ResponseWriter, r *http.Request) {
  var model models.Seller
  var primaryKey = "id"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateSeller(w http.ResponseWriter, r *http.Request) {
  var model models.Seller
  var primaryKey = "id"

  genericUpdate(w, r, &model, primaryKey)
}

func DeleteSeller(w http.ResponseWriter, r *http.Request) {
  var model models.Seller
  var primaryKey = "id"

  genericDelete(w, r, &model, primaryKey)
}
