package models

import (
  "time"
)

type InwardInvoiceDetail struct {
  InwardInvoiceId   int    `gorm:"unique_index"`
  ProductId         int
  Date              time.Time
  Expired           time.Time
  Amount            int
  Price             float64
  Discount1         float64
  Discount2         float64
  Discount3         float64
  Discount4         float64
}
