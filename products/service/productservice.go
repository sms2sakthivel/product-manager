package service

import (
	"github.com/sms2sakthivel/product-manager/products/model"
	"github.com/sms2sakthivel/product-manager/products/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *ProductService) GetProduct(id uint) (*model.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.Repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}
