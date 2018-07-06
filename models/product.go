package models

import (
  "time"
)

type Product struct {
  Id                int     `gorm:"primary_key"`
  Code              string  `gorm:"type:varchar(512)"`
  Name              string  `gorm:"type:varchar(512)"`
  Unit              string  `gorm:"type:varchar(512)"`
  Price             float64
  SpecialPrice      float64
  MinStockWarehouse int
  MinStockDisplay   int
  RackLocation      string  `gorm:"type:varchar(512)"`
  LastModified      time.Time
}
