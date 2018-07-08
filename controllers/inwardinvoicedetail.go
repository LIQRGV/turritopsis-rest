package controllers

import (
  "net/http"
  "turritopsis-rest/models"
)

func ShowInwardInvoiceDetails(w http.ResponseWriter, r *http.Request) {
  GetInwardInvoiceDetail(w,r)
}

func GetInwardInvoiceDetail(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoiceDetail
  var allowedSearchKey = []string{
    "inward_invoice_invoice_code",
    "product_code",
    "date",
    "expired",
  }

  genericQueryGet(w, r, &model, allowedSearchKey)
}

func CreateInwardInvoiceDetail(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoiceDetail
  var primaryKey = ""

  genericCreate(w, r, &model, primaryKey)
}

func UpdateInwardInvoiceDetail(w http.ResponseWriter, r *http.Request) {
  var model models.InwardInvoiceDetail
  var primaryKeys = []string {
    "inward_invoice_invoice_code",
    "product_code",
  }

  genericQueryUpdate(w, r, &model, primaryKeys)
}

