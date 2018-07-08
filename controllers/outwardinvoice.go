package controllers

import (
  "net/http"
  "encoding/json"
  "turritopsis-rest/models"
)

func ShowOutwardInvoices(w http.ResponseWriter, r *http.Request) {
  var models []models.OutwardInvoice

  db.Find(&models)

  json.NewEncoder(w).Encode(models)
}

func GetOutwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.OutwardInvoice
  var primaryKey = "id"

  genericGet(w, r, &model, primaryKey)
}

func CreateOutwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.OutwardInvoice
  var primaryKey = "id"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateOutwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.OutwardInvoice
  var primaryKey = "id"

  genericUpdate(w, r, &model, primaryKey)
}
