package service

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists with the given username or email")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid username or password")

	errFeedbackConsentRequired = errors.New("user consent is required to submit feedback")
	errFeedbackInvalidStatus   = errors.New("invalid feedback status")

// Add other common service-level errors here
)
