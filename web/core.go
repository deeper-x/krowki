package web

import "github.com/kataras/iris"

// CreateApp todo doc
func CreateApp() *iris.Application {
	app := iris.New()

	app.Get("/MooredNow/{id_portinformer:string}", MooredNow)

	return app
}
