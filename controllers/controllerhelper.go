package controllers

import (
  "fmt"
  "os"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func InitController() {
  var dbHost = os.Getenv("DATABASE_HOST")
  var dbDriver = os.Getenv("DATABASE_DRIVER")
  var dbName = os.Getenv("DATABASE_NAME")
  var dbUsername = os.Getenv("DATABASE_USERNAME")
  var dbPassword = os.Getenv("DATABASE_PASSWORD")

  var dbUri = fmt.Sprintf(
    "%s://%s:%s@%s/%s", dbDriver, dbUsername, dbPassword, dbHost, dbName,
  )

  db, err = gorm.Open(
    dbDriver,
    dbUri,
  )
}

