package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"Result": "Hello Iris!"})
	})
	app.Run(iris.Addr(":8080"))
}
