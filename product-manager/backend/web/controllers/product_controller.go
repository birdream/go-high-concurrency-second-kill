package controllers

import (
	"product-manager/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// ProductController controller
type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
}

// GetAll get all product data for view
func (p *ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}
