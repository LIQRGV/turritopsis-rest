package models

import (
  "time"
)

type Warehouse struct {
  ProductCode               string    `gorm:"type:varchar(512);index;not null"`
  InwardInvoiceInvoiceCode  string    `gorm:"type:varchar(512);not null"`
  Amount                    int
  ArrivalTime               time.Time
}

