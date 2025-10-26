package auth

import "time"

type TokenMaker interface {
	CreateToken(username string, role string, duration time.Duration, tokenType TokenType) (string, error)
	VerifyToken(token string) (*Payload, error)
}
