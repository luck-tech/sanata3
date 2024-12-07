package controller

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
)

type SearchSkillTagRequest struct {
	Tag   string `query:"tag"`
	Limit int    `query:"limit"`
}

func (r SearchSkillTagRequest) Validate() error {
	return nil
}

type SearchSkillTagResponse struct {
	Tags []string `json:"tags"`
}

// googleLogin godoc
// @Summary  Serch SkillTag
// @ID       SerchSkillTag
// @Tags     SkillTag
// @Accept   json
// @Produce  json
// @Param 	 b	  query      SearchSkillTagRequest  true "search tag"
// @Success  200  {object}  SearchSkillTagResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/skilltags [get]
func SearchSkillTag(i *interactor.Skill) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var requBody SearchSkillTagRequest
		if err := c.Bind(&requBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := requBody.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, err := i.SearchSkill(ctx, requBody.Tag, requBody.Limit)

		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, SearchSkillTagResponse{
			Tags: result,
		})
	}
}
