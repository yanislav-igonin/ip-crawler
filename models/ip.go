package models

import "time"

type Ip struct {
	Address   string    `gorm:"primaryKey"`
	CheckedAt time.Time `gorm:"column:checkedAt" sql:"DEFAULT:'current_timestamp'"`
	Status    string
}
