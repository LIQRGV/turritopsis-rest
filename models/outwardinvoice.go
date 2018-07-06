package models

import (
  "time"
)

type OutwardInvoice struct {
  Id                int     `gorm:"primary_key"`
  Date              time.Time
  UserId            int
  TotalPrice        float64

  User              User    `gorm:"foreignkey:ID;association_foreignkey:UserId"`
}
