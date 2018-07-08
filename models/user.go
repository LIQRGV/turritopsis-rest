package models

type User struct {
  Id       int
  Username string  `gorm:"type:varchar(512);not null"`
  Password string  `gorm:"type:varchar(512);not null"`
  Role     string  `gorm:"type:varchar(512);not null"`
}
