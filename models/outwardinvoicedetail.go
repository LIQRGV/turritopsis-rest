package models

type OutwardInvoiceDetail struct {
  OutwardInvoiceId          int     `gorm:"unique_index"`
  InwardInvoiceInvoiceCode  string  `gorm:"type:varchar(512)"`
  Price                     float64
  Amount                    int
  TotalPrice                float64
}
