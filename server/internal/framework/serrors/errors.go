package serrors

import "errors"

var (
	ErrSessionNotFound    = errors.New("session not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrPermissionNotFound = errors.New("permission not found")
)
