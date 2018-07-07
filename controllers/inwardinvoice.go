package controllers

import (
  "net/http"
  "encoding/json"
  "turritopsis-rest/models"
)

func ShowInwardInvoices(w http.ResponseWriter, r *http.Request) {
  var inwardInvoice []models.InwardInvoice

  db.Find(&inwardInvoice)

  json.NewEncoder(w).Encode(inwardInvoice)
}

func GetInwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoice
  var primaryKey = "invoice_code"

  genericGet(w, r, &model, primaryKey)
}

func CreateInwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoice
  var primaryKey = "invoice_code"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateInwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoice
  var primaryKey = "invoice_code"

  genericUpdate(w, r, &model, primaryKey)
}

func DeleteInwardInvoice(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoice
  var primaryKey = "invoice_code"

  genericDelete(w, r, &model, primaryKey)
}
