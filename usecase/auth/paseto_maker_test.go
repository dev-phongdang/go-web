package auth

import (
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(paseto.NewV4AsymmetricSecretKey().ExportHex())
	require.NoError(t, err)

	username := "Test-User"
	role := "admin"
	duration := time.Minute * 15

	token, err := maker.CreateToken(username, role, duration, TokenTypeAccessToken)
	require.NoError(t, err)
	require.NotEmpty(t, token)
}
