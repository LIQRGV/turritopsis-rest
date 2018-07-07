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

