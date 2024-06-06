package models

import "github.com/google/uuid"

// LoginRequest model info
// @Description Login request with email and password
type LoginRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
} //@name LoginRequest

// JWT claims
// @Description Created when client calls a endpoint, used for auth and rbac
type JWT struct {
	Bid  uuid.UUID `validate:"required"`
	Role string    `validate:"required"`
} //@name JWT
