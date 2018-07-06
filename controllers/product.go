package controllers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "os"
  "turritopsis-rest/models"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func ShowProducts(w http.ResponseWriter, r *http.Request) {
  var products []models.Product

  var dbHost = os.Getenv("DATABASE_HOST")
  var dbDriver = os.Getenv("DATABASE_DRIVER")
  var dbName = os.Getenv("DATABASE_NAME")
  var dbUsername = os.Getenv("DATABASE_USERNAME")
  var dbPassword = os.Getenv("DATABASE_PASSWORD")

  var dbUri = fmt.Sprintf(
    "%s://%s:%s@%s/%s", dbDriver, dbUsername, dbPassword, dbHost, dbName,
  )

  var db, _ = gorm.Open(
    dbDriver,
    dbUri,
  )

  db.Find(&products)

  json.NewEncoder(w).Encode(products)
}

