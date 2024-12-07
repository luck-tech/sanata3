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

type GetRoomRequest struct {
	RoomID string `param:"roomId"`
}

func (r GetRoomRequest) Validate() error {
	if len(r.RoomID) == 0 {
		return errors.New("roomId is required")
	}
	return nil
}

type GetRoomResponse struct {
	RoomID      string               `json:"roomId"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	OwnerID     string               `json:"ownerId"`
	AimTags     []entity.Skill       `json:"aimTags"`
	Members     []entity.DisplayUser `json:"members"`
}

// googleLogin godoc
// @Summary  Get Room
// @ID       GetRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Param 	 roomId	path 		string  true "room ID"
// @Success  200  {object}  GetRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId} [get]
func GetRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqQuery GetRoomRequest
		if err := c.Bind(&reqQuery); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqQuery.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, err := i.GetByID(ctx, reqQuery.RoomID)
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, GetRoomResponse{
			RoomID:      result.Room.ID,
			Name:        result.Room.Name,
			Description: result.Room.Description,
			OwnerID:     result.Room.OwnerID,
			AimTags:     result.AimTags,
			Members:     result.Members,
		})
	}
}

type ListRoomResponse struct {
	Rooms []GetRoomResponse `json:"rooms"`
}

// googleLogin godoc
// @Summary  List Room
// @ID       ListRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Success  200  {object}  ListRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms [get]
func ListRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)

		result, err := i.List(ctx)
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		var rooms []GetRoomResponse
		for _, room := range result.Rooms {
			rooms = append(rooms, GetRoomResponse{
				RoomID:      room.Room.ID,
				Name:        room.Room.Name,
				Description: room.Room.Description,
				OwnerID:     room.Room.OwnerID,
				AimTags:     room.AimTags,
				Members:     room.Members,
			})
		}

		return c.JSON(http.StatusOK, ListRoomResponse{
			Rooms: rooms,
		})
	}
}

type CreateRoomRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AimSkills   []string `json:"aimSkills"`
	CreatedBy   string   `json:"createdBy"`
}

func (r CreateRoomRequest) Validate() error {
	if len(r.Name) == 0 {
		return errors.New("name is required")
	}

	if len(r.CreatedBy) == 0 {
		return errors.New("created by is required")
	}

	return nil
}

type CreateRoomResponse struct {
	RoomID      string               `json:"roomId"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	OwnerID     string               `json:"ownerId"`
	AimTags     []entity.Skill       `json:"aimTags"`
	Members     []entity.DisplayUser `json:"members"`
}

// googleLogin godoc
// @Summary  Create Room
// @ID       CreateRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Param 	 b			 body 		 CreateRoomRequest  true "create room request"
// @Success  200  {object}  CreateRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms [post]
func CreateRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqBody CreateRoomRequest
		if err := c.Bind(&reqBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqBody.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, err := i.Create(ctx, interactor.CreateRoomParam(reqBody))
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, CreateRoomResponse{
			RoomID:      result.Room.ID,
			Name:        result.Room.Name,
			Description: result.Room.Description,
			OwnerID:     result.Room.OwnerID,
			AimTags:     result.AimTags,
			Members:     result.Members,
		})
	}
}

type UpdateRoomRequest struct {
	RoomID      string   `param:"roomId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AimSkills   []string `json:"aimSkills"`
}

func (r UpdateRoomRequest) Validate() error {
	if len(r.Name) == 0 {
		return errors.New("name is required")
	}

	return nil
}

type UpdateRoomResponse struct {
	RoomID      string               `json:"roomId"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	OwnerID     string               `json:"ownerId"`
	AimTags     []entity.Skill       `json:"aimTags"`
	Members     []entity.DisplayUser `json:"members"`
}

// googleLogin godoc
// @Summary  Update Room
// @ID       UpdateRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Param 	 b			 body 		 UpdateRoomRequest  true "create room request"
// @Success  200  {object}  CreateRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId} [put]
func UpdateRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqBody UpdateRoomRequest
		if err := c.Bind(&reqBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqBody.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		result, err := i.Update(ctx, interactor.UpdateRoomParam{
			RoomID:      reqBody.RoomID,
			Name:        reqBody.Name,
			Description: reqBody.Description,
			AimSkills:   reqBody.AimSkills,
		})
		if err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, UpdateRoomResponse{
			RoomID:      result.Room.ID,
			Name:        result.Room.Name,
			Description: result.Room.Description,
			OwnerID:     result.Room.OwnerID,
			AimTags:     result.AimTags,
		})
	}
}

type DeleteRoomRequest struct {
	RoomID string `param:"roomId"`
}

func (r DeleteRoomRequest) Validate() error {
	return nil
}

type DeleteRoomResponse struct {
	Message string `json:"string"`
}

// googleLogin godoc
// @Summary  Delete Room
// @ID       DeleteRoom
// @Tags     Room
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Success  200  {object}  DeleteRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId} [delete]
func DeleteRoom(i *interactor.Room) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqQuery DeleteRoomRequest
		if err := c.Bind(&reqQuery); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := reqQuery.Validate(); err != nil {
			slog.Error("failed to validate request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := i.Delete(ctx, reqQuery.RoomID); err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, DeleteRoomResponse{Message: "leave sccessful"})
	}
}
