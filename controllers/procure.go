package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/liqrgv/turritopsis-rest/models"
)

func ShowProcures(w http.ResponseWriter, r *http.Request) {
  var models []models.Procure

  db.Find(&models)

  json.NewEncoder(w).Encode(models)
}

func GetProcure(w http.ResponseWriter, r *http.Request) {
  var model models.Procure
  var primaryKey = "id"

  genericGet(w, r, &model, primaryKey)
}

func CreateProcure(w http.ResponseWriter, r *http.Request) {
  var model models.Procure
  var primaryKey = "id"

  genericCreate(w, r, &model, primaryKey)
}

func UpdateProcure(w http.ResponseWriter, r *http.Request) {
  var model models.Procure
  var primaryKey = "id"

  genericUpdate(w, r, &model, primaryKey)
}

func DeleteProcure(w http.ResponseWriter, r *http.Request) {
  var model models.Procure
  var primaryKey = "id"

  genericDelete(w, r, &model, primaryKey)
}
