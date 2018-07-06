package models

type Seller struct {
  Id              int     `gorm:"primary_key"`
  Name            string  `gorm:"type:varchar(512)"`
  Address         string  `gorm:"type:varchar(512)"`
  Phone           string  `gorm:"type:varchar(512)"`
}

