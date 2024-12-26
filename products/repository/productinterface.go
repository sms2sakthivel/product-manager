package repository

import (
	"github.com/sms2sakthivel/product-manager/products/model"
)

type ProductRepository interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	CreateProduct(Product *model.Product) error
	UpdateProduct(Product *model.Product) error
	DeleteProduct(id uint) error
}
