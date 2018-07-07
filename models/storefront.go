package models

import (
  "time"
)

type Storefront struct {
  ProductCode               string     `gorm:"type:varchar(512);index"`
  Amount                    int
  MutationTime              time.Time
  InwardInvoiceInvoiceCode  string     `gorm:"type:varchar(512)"`
}
