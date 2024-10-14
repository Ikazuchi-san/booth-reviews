package models

type Asset struct {
	AssetId     uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	PreviewUrl  string
	Price       float64 // not 32?

	Reviews []Review `gorm:"foreignKey:AssetId"`
	Ratings []Rating `gorm:"foreignKey:AssetId"`
}
