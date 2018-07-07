package models

import (
  "time"
)

type InwardInvoiceDetail struct {
  InwardInvoiceId   int    `gorm:"unique_index"`
  ProductCode       string `gorm:"type:varchar(512);index"`
  Date              time.Time
  Expired           time.Time
  Amount            int
  Price             float64
  Discount1         float64
  Discount2         float64
  Discount3         float64
  Discount4         float64
}
