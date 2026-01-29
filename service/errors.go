package service

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists with the given username or email")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid username or password")

// Add other common service-level errors here
)
