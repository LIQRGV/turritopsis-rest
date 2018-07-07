package main

import (
  "os"
  "fmt"
  "turritopsis-rest/models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type relationStruct struct {
  model     interface{}
  field     string
  dest      string
  onDelete  string
  onUpdate  string
}

var modelCollection = [11]interface{} {
  &models.InwardInvoice{},
  &models.InwardInvoiceDetail{},
  &models.Order{},
  &models.OutwardInvoice{},
  &models.OutwardInvoiceDetail{},
  &models.Procure{},
  &models.Product{},
  &models.Seller{},
  &models.Storefront{},
  &models.User{},
  &models.Warehouse{},
}

var relationCollection = [14]relationStruct {
  relationStruct{
    model: &models.InwardInvoice{},
    field: "procure_id",
    dest: "procures(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.InwardInvoice{},
    field: "user_id",
    dest: "users(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.InwardInvoiceDetail{},
    field: "inward_invoice_id",
    dest: "inward_invoices(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.InwardInvoiceDetail{},
    field: "product_code",
    dest: "products(code)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Order{},
    field: "seller_id",
    dest: "sellers(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Order{},
    field: "user_id",
    dest: "users(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.OutwardInvoice{},
    field: "user_id",
    dest: "users(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.OutwardInvoiceDetail{},
    field: "outward_invoice_id",
    dest: "outward_invoices(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.OutwardInvoiceDetail{},
    field: "inward_invoice_id",
    dest: "inward_invoices(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Procure{},
    field: "seller_id",
    dest: "sellers(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Storefront{},
    field: "product_code",
    dest: "products(code)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Storefront{},
    field: "inward_invoice_id",
    dest: "inward_invoices(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Warehouse{},
    field: "product_code",
    dest: "products(code)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
  relationStruct{
    model: &models.Warehouse{},
    field: "inward_invoice_id",
    dest: "inward_invoices(id)",
    onDelete: "RESTRICT",
    onUpdate: "RESTRICT",
  },
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

  for _, element := range relationCollection {
    db.Model(element.model).AddForeignKey(
      element.field,
      element.dest,
      element.onDelete,
      element.onUpdate,
    )
  }
}
