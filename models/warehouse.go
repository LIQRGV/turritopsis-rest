package models

import (
  "time"
)

type Warehouse struct {
  ProductCode     string    `gorm:"type:varchar(512);index"`
  Amount          int
  ArrivalTime     time.Time
  InwardInvoiceId int
}

