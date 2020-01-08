package main

import (
	"context"

	"product-manager/backend/web/controllers"
	"product-manager/common"
	"product-manager/repositories"
	"product-manager/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/prometheus/common/log"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	template := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").Reload(true)

	app.RegisterView(template)

	app.StaticWeb("/assets", "./web/assets")

	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 5.注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	orderRepository := repositories.NewOrderMangerRepository("order", db)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))

	//6.启动服务
	app.Run(
		iris.Addr("localhost:8081"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
