package controllers

import (
  "net/http"
  "github.com/liqrgv/turritopsis-rest/models"
)

func ShowOutwardInvoiceDetails(w http.ResponseWriter, r *http.Request) {
  GetOutwardInvoiceDetail(w,r)
}

func GetOutwardInvoiceDetail(w http.ResponseWriter, r *http.Request) {
  var model models.OutwardInvoiceDetail
  var allowedSearchKey = []string{
    "inward_invoice_invoice_code",
    "outward_invoice_id",
  }

  genericQueryGet(w, r, &model, allowedSearchKey)
}

func CreateOutwardInvoiceDetail(w http.ResponseWriter, r *http.Request) {
  var model models.OutwardInvoiceDetail
  var primaryKey = ""

  genericCreate(w, r, &model, primaryKey)
}

func UpdateOutwardInvoiceDetail(w http.ResponseWriter, r *http.Request) {
  var model models.OutwardInvoiceDetail
  var primaryKeys = []string {
    "inward_invoice_invoice_code",
    "outward_invoice_id",
  }

  genericQueryUpdate(w, r, &model, primaryKeys)
}

