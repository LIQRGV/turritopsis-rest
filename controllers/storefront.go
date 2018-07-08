package controllers

import (
  "net/http"
  "turritopsis-rest/models"
)

func ShowStorefronts(w http.ResponseWriter, r *http.Request) {
  GetStorefront(w,r)
}

func GetStorefront(w http.ResponseWriter, r *http.Request) {
  var model models.Storefront
  var allowedSearchKey = []string{
    "product_code",
    "inward_invoice_invoice_code",
  }

  genericQueryGet(w, r, &model, allowedSearchKey)
}

func CreateStorefront(w http.ResponseWriter, r *http.Request) {
  var model models.Storefront
  var primaryKey = ""

  genericCreate(w, r, &model, primaryKey)
}

func UpdateStorefront(w http.ResponseWriter, r *http.Request) {
  var model models.Storefront
  var primaryKeys = []string {
    "product_code",
    "inward_invoice_invoice_code",
  }

  genericQueryUpdate(w, r, &model, primaryKeys)
}

