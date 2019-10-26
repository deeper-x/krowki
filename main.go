package main

import (
	"github.com/deeper-x/krowki/web"
	"github.com/kataras/iris"
)

func main() {
	app := web.CreateApp()
	app.Run(iris.Addr(":8000"))
}
