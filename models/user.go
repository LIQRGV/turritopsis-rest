package models

type User struct {
  Id       int
  Username string  `gorm:"type:varchar(512)"`
  Password string  `gorm:"type:varchar(512)"`
  Role     string  `gorm:"type:varchar(512)"`
}
