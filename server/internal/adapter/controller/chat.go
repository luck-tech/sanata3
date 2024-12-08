package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
)

// googleLogin godoc
// @Summary  Get Chat
// @ID       GetChat
// @Tags     Chat
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Success  200  {object}  JoinRoomResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId}/chat [get]
func JoinChatRoom(i *interactor.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		roomId := c.Param("roomId")

		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ctx := contexts.ConvertContext(c)

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		var lastTime time.Time

		for {
			select {
			case <-c.Request().Context().Done():
				slog.Info("SSE client disconnect")
				return nil
			case <-ticker.C:
				chats, err := i.Get(ctx, interactor.GetChatParam{
					RoomID:   roomId,
					LastTime: lastTime,
				})
				if err != nil {
					slog.Error("failed to get chat", "error", err)
					if err := NewEvent(map[string]string{"error": err.Error()}).MarshalTo(w); err != nil {
						slog.Error("failed to write event", "error", err)
					}
				} else {
					if chats != nil {
						lastTime = time.Now()
					}
					if err := NewEvent(chats).MarshalTo(w); err != nil {
						slog.Error("failed to write event", "error", err)
					}
				}
				w.Flush()
			}
		}
	}
}

type Event struct {
	Data    any
	Retry   []byte
	Comment []byte
}

func NewEvent(data any) *Event {
	if data == nil {
		return &Event{
			Comment: []byte("ping"),
		}
	}
	return &Event{
		Data: data,
	}
}

func (e *Event) MarshalTo(w io.Writer) error {
	if e.Data == nil && len(e.Comment) == 0 {
		return nil
	}

	if e.Data != nil {
		data, err := json.Marshal(e.Data)
		if err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "%s", data); err != nil {
			return err
		}

		if len(e.Retry) > 0 {
			if _, err := fmt.Fprintf(w, "retry: %s\n", e.Retry); err != nil {
				return err
			}
		}
	}

	if len(e.Comment) > 0 {
		if _, err := fmt.Fprintf(w, ": %s\n", e.Comment); err != nil {
			return err
		}
	}

	return nil
}

type PostChatRequest struct {
	RoomID  string `param:"roomId"`
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

type PostChatResponse struct {
	Message string `json:"string"`
}

// googleLogin godoc
// @Summary  Post Chat
// @ID       PostChat
// @Tags     Chat
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Param 	 b		body 		 PostChatRequest  true "post chat request"
// @Success  200  {object}  PostChatResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId}/chat [post]
func PostChat(i *interactor.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqBody PostChatRequest
		if err := c.Bind(&reqBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := i.Post(ctx, interactor.PostChatParam{
			RoomID:  reqBody.RoomID,
			Message: reqBody.Message,
		}); err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, PostChatResponse{Message: "post sccessful"})
	}
}

type EditChatRequest struct {
	RoomID  string `param:"roomId"`
	ChatID  string `param:"chatId"`
	UserID  string `json:"userId"`
	Message string `json:"message"`
}

type EditChatResponse struct {
	Message string `json:"string"`
}

// googleLogin godoc
// @Summary  Edit Chat
// @ID       EditChat
// @Tags     Chat
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Param 	 chatId		path 		 string  true "chatID path param"
// @Param 	 b		body 		 EditChatRequest  true "edit chat request"
// @Success  200  {object}  EditChatResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId}/chat/{chatId} [put]
func EditChat(i *interactor.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqBody EditChatRequest
		if err := c.Bind(&reqBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := i.Edit(ctx, interactor.UpdateChatParam{
			RoomID:  reqBody.RoomID,
			ChatID:  reqBody.ChatID,
			Message: reqBody.Message,
		}); err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, EditChatResponse{Message: "post sccessful"})
	}
}

type DeleteChatRequest struct {
	RoomID string `param:"roomId"`
	ChatID string `param:"chatId"`
}

type DeleteChatResponse struct {
	Message string `json:"string"`
}

// googleLogin godoc
// @Summary  Delete Chat
// @ID       DeleteChat
// @Tags     Chat
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Param 	 chatId		path 		 string  true "chatID path param"
// @Success  200  {object}  DeleteChatResponse
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId}/chat/{chatId} [delete]
func DeleteChat(i *interactor.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := contexts.ConvertContext(c)
		var reqBody DeleteChatRequest
		if err := c.Bind(&reqBody); err != nil {
			slog.Error("failed to bind request body", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrBadRequest
		}

		if err := i.Delete(ctx, interactor.DeleteChatParam{
			RoomID: reqBody.RoomID,
			ChatID: reqBody.ChatID,
		}); err != nil {
			slog.Error("failed to login github", "error", err, "requestID", contexts.GetRequestID(ctx))
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, DeleteChatResponse{Message: "post sccessful"})
	}
}
