package models

import (
  "time"
)

type InwardInvoice struct {
  InvoiceCode   string  `gorm:"type:varchar(512);primary_key"`
  ProcureId     int
  Date          time.Time
  TotalPrice    float64
  DueDate       time.Time
  UserId        int
  Discount      float64
  Tax           float64
}
