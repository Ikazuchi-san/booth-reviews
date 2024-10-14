package models

import "time"

type Rating struct {
	RatingId    uint      `gorm:"primaryKey"`
	AssetId     uint      `gorm:"not null"`
	UserId      uint      `gorm:"not null"`
	RatingValue int       `gorm:"not null;range:1,5"`
	Timestamp   time.Time `gorm:"autoCreateTime"`
}
