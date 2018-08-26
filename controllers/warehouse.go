package controllers

import (
  "net/http"
  "github.com/liqrgv/turritopsis-rest/models"
)

func ShowWarehouses(w http.ResponseWriter, r *http.Request) {
  GetWarehouse(w,r)
}

func GetWarehouse(w http.ResponseWriter, r *http.Request) {
  var model models.Warehouse
  var allowedSearchKey = []string{
    "product_code",
    "inward_invoice_invoice_code",
  }

  genericQueryGet(w, r, &model, allowedSearchKey)
}

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
  var model models.Warehouse
  var primaryKey = ""

  genericCreate(w, r, &model, primaryKey)
}

func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
  var model models.Warehouse
  var primaryKeys = []string {
    "product_code",
    "inward_invoice_invoice_code",
  }

  genericQueryUpdate(w, r, &model, primaryKeys)
}

