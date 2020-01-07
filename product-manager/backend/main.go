package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	template := iris.HTML("./backend/web/views", ".html").
		Layout("shared/layout.html").Reload(true)

	app.RegisterView(template)

	app.StaticWeb("/assets", "./backend/web/assets")

	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//6.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
