package models

import (
  "time"
)

type InwardInvoice struct {
  Id            int     `gorm:"primary_key"`
  ProcureId     int
  InvoiceCode   string  `gorm:"type:varchar(512)"`
  Date          time.Time
  TotalPrice    float64
  DueDate       time.Time
  UserId        int
  Discount      float64
  Tax           float64
}
