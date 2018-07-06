package models

import (
  "time"
)

type Order struct {
  Id                int     `gorm:"primary_key"`
  Date              time.Time
  SellerId          int
  UserId            int
}
