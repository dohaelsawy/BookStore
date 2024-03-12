package token

import (
	"testing"
	"time"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomString(32)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username,duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.Equal(t, "7", payload.RegisteredClaims.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.RegisteredClaims.IssuedAt.Time, time.Minute)
	require.WithinDuration(t, expiredAt, payload.RegisteredClaims.ExpiresAt.Time, time.Minute)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomString(32), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}