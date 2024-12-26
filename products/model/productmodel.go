package model

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Brand       string `gorm:"size:100;not null"`
	Category    string `gorm:"size:100;not null"`
	SubCategory string `gorm:"size:100;not null"`
	Price       uint   `gorm:"not null"`
}

func (product *Product) GetAPIResponseObject() *ProductResponse {
	return &ProductResponse{ID: product.ID, Name: product.Name, Brand: product.Brand, Category: product.Category, SubCategory: product.SubCategory, Price: product.Price}
}
