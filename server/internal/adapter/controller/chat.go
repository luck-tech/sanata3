package controller

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
	"golang.org/x/net/websocket"
)

// googleLogin godoc
// @Summary  Get Chat
// @ID       GetChat
// @Tags     Chat
// @Accept   json
// @Produce  json
// @Param 	 roomId		path 		 string  true "roomID path param"
// @Success  200  {array}  entity.Chat
// @Failure  400  {object}  echo.HTTPError
// @Failure  500  {object}  echo.HTTPError
// @Router   /v1/rooms/{roomId}/chat [get]
func JoinChatRoom(i *interactor.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		roomId := c.Param("roomId")

		ctx := contexts.ConvertContext(c)

		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			ticker := time.NewTicker(1 * time.Second)
			var lastTime = time.Time{}
			defer ticker.Stop()

			for {
				select {
				case <-c.Request().Context().Done():
					return
				case <-ticker.C:
					chats, err := i.Get(ctx, interactor.GetChatParam{
						RoomID:   roomId,
						LastTime: lastTime,
					})
					log.Println("ちゃっと！", chats)

					if err != nil {
						slog.Error("failed to get chat", "error", err)
						data, err := MarshalTo(map[string]string{"error": err.Error()})
						if err != nil {
							slog.Error("failed to marshal event", "error", err)
						}

						if err := websocket.Message.Send(ws, string(data)); err != nil {
							c.Logger().Error(err)
						}
					} else {
						if len(chats) != 0 {
							lastTime = time.Now()
							data, err := MarshalTo(chats)
							if err != nil {
								slog.Error("failed to marshal event", "error", err)
							}
							if err := websocket.Message.Send(ws, string(data)); err != nil {
								c.Logger().Error(err)
							}
						}
					}
				}
			}
		}).ServeHTTP(c.Response(), c.Request())

		return nil
	}
}

func MarshalTo(data any) ([]byte, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return d, nil
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
