package web

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestWeb(t *testing.T) {
	app := CreateApp()

	e := httptest.New(t, app)

	e.GET("/MooredNow/28").Expect().Status(httptest.StatusOK)
}
