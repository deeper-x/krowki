package web

import "github.com/kataras/iris"

// CreateApp todo doc
func CreateApp() *iris.Application {
	app := iris.New()

	app.Get("/mooredNow/{id_portinformer:string}", MooredNow)
	app.Get("/anchoredNow/{id_portinformer:string}", AnchoredNow)
	app.Get("/arrivalPrevisions/{id_portinformer:string}", AllArrivalPrevisions)
	app.Get("/arrivals/{id_portinformer:string}", AllArrivals)

	return app
}
