package models

import (
  "time"
)

type Procure struct {
  Id              int     `gorm:"primary_key"`
  Date            time.Time
  SellerId        int
  UserId          int

  InwardInvoice   InwardInvoice
}
