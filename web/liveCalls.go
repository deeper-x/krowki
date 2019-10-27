package web

import (
	"log"

	"github.com/deeper-x/krowki/ldb"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
)

// MooredNow todo doc
func MooredNow(ctx iris.Context) {
	db, err := ldb.Connect()

	if err != nil {
		log.Println(err)
	}

	conn := ldb.NewConnection(db)
	idPortinformer := ctx.Params().Get("id_portinformer")
	allMoored := conn.AllMoored(idPortinformer)

	ctx.JSON(allMoored)
}

// AnchoredNow todo doc
func AnchoredNow(ctx iris.Context) {
	db, err := ldb.Connect()

	if err != nil {
		log.Println(err)
	}

	conn := ldb.NewConnection(db)
	idPortinformer := ctx.Params().Get("id_portinformer")
	allAnchored := conn.AllAnchored(idPortinformer)

	ctx.JSON(allAnchored)
}

// AllArrivalPrevisions todo doc
func AllArrivalPrevisions(ctx iris.Context) {
	db, err := ldb.Connect()

	if err != nil {
		log.Println(err)
	}

	conn := ldb.NewConnection(db)
	idPortinformer := ctx.Params().Get("id_portinformer")
	allArrivalPrevisions := conn.AllArrivalPrevisions(idPortinformer)

	ctx.JSON(allArrivalPrevisions)
}

// AllArrivals todo doc
func AllArrivals(ctx iris.Context) {
	db, err := ldb.Connect()

	if err != nil {
		log.Println(err)
	}

	conn := ldb.NewConnection(db)
	idPortinformer := ctx.Params().Get("id_portinformer")
	allArrivals := conn.GetTodayArrivals(idPortinformer)

	ctx.JSON(allArrivals)
}
