package repository

import (
	"github.com/sms2sakthivel/product-manager/products/model"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	DB *gorm.DB
}

func (repo *GormProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := repo.DB.Find(&products).Error
	return products, err
}

func (repo *GormProductRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	err := repo.DB.First(&product, id).Error
	return &product, err
}

func (repo *GormProductRepository) CreateProduct(product *model.Product) error {
	return repo.DB.Create(product).Error
}

func (repo *GormProductRepository) UpdateProduct(product *model.Product) error {
	return repo.DB.Save(product).Error
}

func (repo *GormProductRepository) DeleteProduct(id uint) error {
	return repo.DB.Delete(&model.Product{}, id).Error
}
