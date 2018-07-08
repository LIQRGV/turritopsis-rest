package models

type OutwardInvoiceDetail struct {
  OutwardInvoiceId          int     `gorm:"unique_index;not null"`
  InwardInvoiceInvoiceCode  string  `gorm:"type:varchar(512);not null"`
  Price                     float64
  Amount                    int
  TotalPrice                float64
}
