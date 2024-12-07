package entity

import "time"

type Chat struct {
	ID       string `json:"id" dynamodbav:"id"`
	RoomID   string `json:"roomId" dynamodbav:"roomId"`
	Message  string `json:"message" dynamodbav:"message"`
	AutherID string `json:"autherId" dynamodbav:"autherId"`

	CreatedAt time.Time  `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" dynamodbav:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" dynamodbav:"deletedAt,omitempty"`
}
