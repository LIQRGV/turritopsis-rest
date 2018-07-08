package models

import (
  "time"
)

type Storefront struct {
  ProductCode               string     `gorm:"type:varchar(512);index;not null"`
  InwardInvoiceInvoiceCode  string     `gorm:"type:varchar(512);not null"`
  Amount                    int
  MutationTime              time.Time
}
