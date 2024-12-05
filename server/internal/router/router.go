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

	loginRoute := engine.Group("/login")
	{
		loginRoute.POST("/github", controller.LoginGitHub(interactors.login))
	}

	v1Route := engine.Group("/v1")
	v1Route.Use(middleware.Auth(interactors.login))

	{
		usersRoute := v1Route.Group("/users")
		{
			userRoute := usersRoute.Group("/:userId")
			{
				userRoute.GET("", controller.GetUser(interactors.user))    // figma: profile
				userRoute.PUT("", controller.UpdateUser(interactors.user)) // figma: form
			}
		}

		roomsRoute := v1Route.Group("/rooms")
		{
			roomsRoute.GET("", nil)  // figma: home
			roomsRoute.POST("", nil) // figma: room-create

			roomRoute := roomsRoute.Group("/:roomId")
			{
				roomRoute.GET("", nil)    // figma: room-description
				roomRoute.PUT("", nil)    // figma: room-description
				roomRoute.DELETE("", nil) // figma: room-description

				roomRoute.POST("/members ", nil)   // figma: room-description
				roomRoute.DELETE("/members ", nil) // figma: room-description

				roomRoute.GET("/chat", nil)  // figma: room
				roomRoute.POST("/chat", nil) // figma: room
			}
		}

		recommendsRoute := v1Route.Group("/recommends")
		{
			recommendsRoute.GET("/rooms", nil) // figma: home
			recommendsRoute.GET("/users", nil) // figma: home
		}

		v1Route.GET("/search", nil)                                              // figma: search
		v1Route.GET("/skilltags ", controller.SearchSkillTag(interactors.skill)) // figma: search
	}

	return engine
}
