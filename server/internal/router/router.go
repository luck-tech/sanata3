package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/adapter/controller"
	"github.com/murasame29/go-httpserver-template/internal/adapter/middleware"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// NewEcho は、echo/v4 を利用した http.Handlerを返す関数です。
func NewEcho(interactors *di) http.Handler {
	engine := echo.New()

	engine.GET("/healthz", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	engine.Use(
		middleware.RequestID(),
		echoMiddleware.CORS(),
	)

	loginRoute := engine.Group("/login")
	{
		loginRoute.POST("/github", controller.LoginGitHub(interactors.login))
	}

	v1Route := engine.Group("/v1")
	v1Route.Use(
		echoMiddleware.CORS(),
		middleware.RequestID(),
		middleware.Auth(interactors.login),
	)

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
			roomsRoute.GET("", controller.ListRoom(interactors.room))    // figma: home
			roomsRoute.POST("", controller.CreateRoom(interactors.room)) // figma: room-create

			roomRoute := roomsRoute.Group("/:roomId")
			{
				roomRoute.GET("", controller.GetRoom(interactors.room))       // figma: room-description
				roomRoute.PUT("", controller.UpdateRoom(interactors.room))    // figma: room-description
				roomRoute.DELETE("", controller.DeleteRoom(interactors.room)) // figma: room-description

				roomRoute.POST("/members ", controller.JoinRoom(interactors.room))    // figma: room-description
				roomRoute.DELETE("/members ", controller.LeaveRoom(interactors.room)) // figma: room-description

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
		v1Route.GET("/skilltags", controller.SearchSkillTag(interactors.skill)) // figma: search
	}

	return engine
}
