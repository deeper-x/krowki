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
