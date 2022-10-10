package models

import "time"

type Ip struct {
	Address    string    `gorm:"primaryKey"`
	CheckedAt  time.Time `gorm:"column:checkedAt; not null"`
	Status     string    `gorm:"index; not null"`
	StatusCode int       `gorm:"column:statusCode; index; not null"`
}
