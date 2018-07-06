package models

import (
  "time"
)

type Warehouse struct {
  ProductId       int     `gorm:"index"`
  Amount          int
  ArrivalTime     time.Time
  InwardInvoiceId int
}

