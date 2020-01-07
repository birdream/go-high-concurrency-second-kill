package services

import (
	"product-manager/datamodels"
	"product-manager/repositories"
)

// IProductService interface
type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
}

// ProductService real struct
type ProductService struct {
	productRepository repositories.IProduct
}

// NewProductService init method
func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{repository}
}

// GetProductByID ..
func (p *ProductService) GetProductByID(productID int64) (*datamodels.Product, error) {
	return p.productRepository.SelectByKey(productID)
}

// GetAllProduct ..
func (p *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

// DeleteProductByID ..
func (p *ProductService) DeleteProductByID(productID int64) bool {
	return p.productRepository.Delete(productID)
}

// InsertProduct ..
func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

// UpdateProduct ..
func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}
