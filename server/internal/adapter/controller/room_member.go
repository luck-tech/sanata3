package controller

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
)

type JoinRoomRequest struct {
	RoomID string `param:"roomId"`
}

func (r JoinRoomRequest) Validate() error {
	return nil
}

type JoinRoomResponse struct {
	GetRoomResponse
}

// googleLogin godoc
// @Summary  Join Room
// @ID       JoinRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Success  200  {object}  JoinRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /rooms/{roomId}/members [post]
func JoinRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqQuery JoinRoomRequest
		if err := c.Bind(&reqQuery); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, err := i.Join(ctx, reqQuery.RoomID)
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, JoinRoomResponse{
			GetRoomResponse{
				RoomID:      result.Room.ID,
				Description: result.Room.Description,
				OwnerID:     result.Room.OwnerID,
				AimTags:     result.AimTags,
			},
		})
	}
}

type LeaveRoomRequest struct {
	RoomID string `param:"roomId"`
}

func (r LeaveRoomRequest) Validate() error {
	return nil
}

type LeaveRoomResponse struct {
	Message string `json:"string"`
}

// googleLogin godoc
// @Summary  Leave Room
// @ID       LeaveRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Success  200  {object}  LeaveRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /rooms/{roomId}/members [delete]
func LeaveRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqQuery LeaveRoomRequest
		if err := c.Bind(&reqQuery); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqQuery.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := i.Leave(ctx, reqQuery.RoomID); err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, LeaveRoomResponse{Message: "leave sccessful"})
	}
}
