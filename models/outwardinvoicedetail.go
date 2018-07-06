package models

type OutwardInvoiceDetail struct {
  OutwardInvoiceId  int     `gorm:"unique_index"`
  BatchId           int
  Price             float64
  Amount            int
  TotalPrice        float64
}
