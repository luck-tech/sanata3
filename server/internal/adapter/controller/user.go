package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
)

type GetUserRequest struct {
	UserID string `param:"userId"`
}

func (r GetUserRequest) Validate() error {
	if len(r.UserID) == 0 {
		return errors.New("userId is required")
	}

	if r.UserID == "undefined" {
		return errors.New("userId is required")
	}

	return nil
}

type GetUserResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
	Description string `josn:"description"`

	WantLeanSkills []entity.Skill `json:"wantLeanSkills"`
	UsedSkills     []entity.Skill `json:"usedSkills"`
}

// googleLogin godoc
// @Summary  Get User
// @ID       GetUser
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param 	 userId	 path 		 string  true "userID"
// @Success  200  {object}  GetUserResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /users/{userId} [get]
func GetUser(i *interactor.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqQuery GetUserRequest
		if err := c.Bind(&reqQuery); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqQuery.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, used, want, err := i.GetUser(ctx, reqQuery.UserID)
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, GetUserResponse{
			ID:             result.ID,
			Name:           result.Name,
			Email:          result.Email,
			Icon:           result.Icon,
			Description:    result.Description,
			WantLeanSkills: want,
			UsedSkills:     used,
		})
	}
}

type UpdateUserRequest struct {
	UserID          string   `param:"userId"`
	Description     string   `json:"description"`
	WantLearnSkills []string `json:"wantLearnSkills"`
	UsedSkills      []string `json:"usedSkills"`
}

func (r UpdateUserRequest) Validate() error {
	if len(r.UserID) == 0 {
		return errors.New("userId is required")
	}

	if r.UserID == "undefined" {
		return errors.New("userId is required")
	}

	return nil
}

type UpdateUserResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
	Description string `josn:"description"`

	WantLeanSkills []entity.Skill `json:"wantLeanSkills"`
	UsedSkills     []entity.Skill `json:"usedSkills"`
}

// googleLogin godoc
// @Summary  Update User
// @ID       UpdateUser
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param 	 userId	 path 		  string             true "userID"
// @Param 	 b	     body 		  UpdateUserRequest  true "update user request"
// @Success  200  {object}  UpdateUserResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /users/{userId} [put]
func UpdateUser(i *interactor.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var requBody UpdateUserRequest
		if err := c.Bind(&requBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := requBody.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, used, want, err := i.UpdateUser(ctx, interactor.UpdateUserParam{
			UserID:          requBody.UserID,
			Description:     requBody.Description,
			WantLearnSkills: requBody.WantLearnSkills,
			UsedSkills:      requBody.UsedSkills,
		})
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, GetUserResponse{
			ID:             result.ID,
			Name:           result.Name,
			Email:          result.Email,
			Icon:           result.Icon,
			Description:    result.Description,
			WantLeanSkills: want,
			UsedSkills:     used,
		})
	}
}
