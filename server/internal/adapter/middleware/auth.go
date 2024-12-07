package middleware

import (
	"errors"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/framework/serrors"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
)

func Auth(login *interactor.Login) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if len(authHeader) == 0 {
				slog.Error("auth error. required auth header")
				return echo.ErrUnauthorized
			}

			ctx := contexts.ConvertContext(c)

			result, err := login.CheckLogin(ctx, authHeader)
			if err != nil {
				if errors.Is(err, serrors.ErrSessionNotFound) {
					slog.Error("auth error. please login", "error", err)
					return echo.ErrUnauthorized
				}
				slog.Error("auth error. check login error", "error", err)
				return echo.ErrInternalServerError
			}

			c.Set(contexts.UserID.String(), result.UserID)
			c.Set(contexts.SessionID.String(), result.SessionID)

			return next(c)
		}
	}
}
