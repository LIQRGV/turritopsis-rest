package models

import (
  "time"
)

type StoreFront struct {
  ProductId       int     `gorm:"index"`
  Amount          int
  MutationTime    time.Time
  InwardInvoiceId int
}
