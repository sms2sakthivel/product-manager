package service

import (
	"github.com/sms2sakthivel/product-manager/products/model"
	"github.com/sms2sakthivel/product-manager/products/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (s *ProductService) GetProducts() ([]model.ProductResponse, error) {
	var products []model.ProductResponse = []model.ProductResponse{}
	dbProducts, err := s.Repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	for _, product := range dbProducts {
		products = append(products, *product.GetAPIResponseObject())
	}
	return products, nil
}

func (s *ProductService) GetProduct(id uint) (*model.ProductResponse, error) {
	dbProduct, err := s.Repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return dbProduct.GetAPIResponseObject(), nil
}

func (s *ProductService) CreateProduct(productReq *model.ProductCreateRequest) (*model.ProductResponse, error) {
	product := productReq.GetDBObject()
	err := s.Repo.CreateProduct(product)
	return product.GetAPIResponseObject(), err
}

func (s *ProductService) UpdateProduct(productReq *model.ProductUpdateRequest) (*model.ProductResponse, error) {
	product := productReq.GetDBObject()
	err := s.Repo.UpdateProduct(product)
	return product.GetAPIResponseObject(), err
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}
