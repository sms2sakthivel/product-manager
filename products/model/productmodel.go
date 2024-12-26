package model

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Brand       string `gorm:"size:100;not null"`
	Category    string `gorm:"size:100;not null"`
	SubCategory string `gorm:"size:100;not null"`
}
