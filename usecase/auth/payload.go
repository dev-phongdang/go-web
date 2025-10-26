package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type TokenType int

const (
	TokenTypeAccessToken  = 1
	TokenTypeRefreshToken = 2
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Type      TokenType `json:"token_type"`
	Username  string    `json:"username"`
	Scope     string    `json:"scope"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, scope string, tokenType TokenType, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Type:      tokenType,
		Username:  username,
		Scope:     scope,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid(tokenType TokenType) error {
	if payload.Type != tokenType {
		return ErrInvalidToken
	}
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
