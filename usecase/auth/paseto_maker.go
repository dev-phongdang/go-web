package auth

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	asymmetric_key string
}

func NewPasetoMaker(asymmetric_key string) (TokenMaker, error) {
	return &PasetoMaker{
		asymmetric_key: asymmetric_key,
	}, nil
}

func (m *PasetoMaker) CreateToken(username string, scope string, duration time.Duration, tokenType TokenType) (string, error) {
	payload, err := NewPayload(username, scope, tokenType, duration)
	if err != nil {
		return "", err
	}

	// Convert Payload struct to a map[string]interface{} for paseto.MakeToken
	// This is necessary because paseto.MakeToken expects a map for claims.
	claims := make(map[string]any)
	claims["id"] = payload.ID
	claims["token_type"] = payload.Type
	claims["username"] = payload.Username
	claims["scope"] = payload.Scope
	claims["issued_at"] = payload.IssuedAt
	claims["expired_at"] = payload.ExpiredAt

	token, err := paseto.MakeToken(claims, []byte{})
	if err != nil {
		return "", err
	}
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(m.asymmetric_key)
	if err != nil {
		return "", err
	}
	// log.Println(secretKey) // don't share this!!!
	// publicKey := secretKey.Public()
	// log.Println(publicKey)
	signed := token.V4Sign(secretKey, nil)
	return signed, err
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	return nil, nil
}
