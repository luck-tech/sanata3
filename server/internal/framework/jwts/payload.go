package jwts

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// token が返すエラー
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// トークン内のペイロードに含まれるデータ
type Payload struct {
	ID        uuid.UUID `json:"id"`
	SessionID string    `json:"sessionID"`
	IssuedAt  time.Time `josn:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

// 新たなペイロードを作る
func NewPayload(sessionID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		SessionID: sessionID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// トークンペイロードを検証する
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
