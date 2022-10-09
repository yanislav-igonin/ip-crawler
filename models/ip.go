package models

import (
	"time"

	"gorm.io/gorm"
)

type Ip struct {
	gorm.Model
	Address   string
	CheckedAt time.Time
	Status    string
}
