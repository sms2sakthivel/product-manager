package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Brand       string `gorm:"not null"`
	Category    string `gorm:"not null"`
	SubCategory string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
}

func (product *Product) GetAPIResponseObject() *ProductResponse {
	return &ProductResponse{ID: product.ID, Name: product.Name, Brand: product.Brand, Category: product.Category, SubCategory: product.SubCategory, Price: product.Price}
}
