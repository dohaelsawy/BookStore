package token

import (
	"testing"
	"time"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
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

	require.Equal(t, "1", payload.RegisteredClaims.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.RegisteredClaims.IssuedAt.Time, time.Minute)
	require.WithinDuration(t, expiredAt, payload.RegisteredClaims.ExpiresAt.Time, time.Minute)
}


func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomString(32), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}