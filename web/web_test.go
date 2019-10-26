package web

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestWeb(t *testing.T) {
	app := CreateApp()

	e := httptest.New(t, app)

	e.GET("/mooredNow/28").Expect().Status(httptest.StatusOK)
	e.GET("/anchoredNow/28").Expect().Status(httptest.StatusOK)

}
