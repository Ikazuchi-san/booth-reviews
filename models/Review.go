package models

import "time"

type Review struct {
	ReviewId   uint `gorm:"primaryKey"`
	AssetId    uint `gorm:"not null"`
	UserId     uint `gorm:"not null"`
	ReviewText uint
	Timestamp  time.Time `gorm:"autoCreateTime"`
}
