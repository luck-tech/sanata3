package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/adapter/controller"
	"github.com/murasame29/go-httpserver-template/internal/adapter/middleware"
)

// NewEcho は、echo/v4 を利用した http.Handlerを返す関数です。
func NewEcho(interactors *di) http.Handler {
	engine := echo.New()

	engine.GET("/healthz", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	engine.Use(
		middleware.RequestID(),
		middleware.SetupCORS(),
	)

	// GitHub Login
	engine.POST("/login/github", controller.LoginGitHub(interactors.login))

	return engine
}
