package controllers

import (
	"product-manager/common"
	"product-manager/datamodels"
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

// PostUpdate 修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()

	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	p.Ctx.Redirect("/product/all")
}
