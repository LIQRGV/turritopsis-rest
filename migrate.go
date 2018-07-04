package main

import (
  "os"
  "fmt"
  "turritopsis-rest/models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var modelCollection = [1]interface{}{
  &models.User{},
}

func migrate() {
  var dbHost = os.Getenv("DATABASE_HOST")
  var dbDriver = os.Getenv("DATABASE_DRIVER")
  var dbName = os.Getenv("DATABASE_NAME")
  var dbUsername = os.Getenv("DATABASE_USERNAME")
  var dbPassword = os.Getenv("DATABASE_PASSWORD")

  var dbUri = fmt.Sprintf(
    "%s://%s:%s@%s/%s", dbDriver, dbUsername, dbPassword, dbHost, dbName,
  )

  var db, err = gorm.Open(
    dbDriver,
    dbUri,
  )

  if(err != nil) {
    fmt.Println(err)
  }

  for _, element := range modelCollection {
    db.AutoMigrate(element)
  }
}
