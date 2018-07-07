package models

import (
  "time"
)

// we use varchar on code,
// because it's possible that code contain leading zero
type Product struct {
  Id                int     `gorm:"primary_key"`
  Code              string  `gorm:"type:varchar(512);unique_index"`
  Name              string  `gorm:"type:varchar(512)"`
  Unit              string  `gorm:"type:varchar(512)"`
  Price             float64
  SpecialPrice      float64
  MinStockWarehouse int
  MinStockDisplay   int
  RackLocation      string  `gorm:"type:varchar(512)"`
  LastModified      time.Time
}
