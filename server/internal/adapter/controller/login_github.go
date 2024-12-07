package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
)

type LoginGitHubRequest struct {
	Code string `json:"code"`
}

func (r LoginGitHubRequest) Validate() error {
	if len(r.Code) == 0 {
		return errors.New("code is required")
	}

	return nil
}

type LoginGitHubResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Code      string `json:"code"`
	IsNewUser bool   `json:"isNewUser"`
}

// googleLogin godoc
// @Summary  GitHub Login
// @ID       LoginGitHub
// @Tags     LoginRequest
// @Accept   json
// @Produce  json
// @Param 	 q			 body 		 LoginGitHubRequest  true "LoginGitHubRequest"
// @Success  200  	 {object}  LoginGitHubResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /login/github [post]
func LoginGitHub(i *interactor.Login) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqBody LoginGitHubRequest
		if err := c.Bind(&reqBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqBody.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, err := i.GitHub(ctx, interactor.LoginGitHubParam{
			Code: reqBody.Code,
		})
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, LoginGitHubResponse{
			ID:        result.UserID,
			Name:      result.UserName,
			Icon:      result.Icon,
			Code:      result.JWT,
			IsNewUser: result.IsNewUser,
		})
	}
}
