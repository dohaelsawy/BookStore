package token

import (
	"errors"
	"strconv"
	"time"

	"github.com/dohaelsawy/bookStore/util"
	"github.com/golang-jwt/jwt/v5"
	//uuid "github.com/google/uuid"
)

type Payload struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID:= util.RandomNumber(1,60)
	payload := &Payload{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID: strconv.FormatInt(int64(tokenID), 10),
		},
	}
	return payload, nil
}


func (payload *Payload) Valid() error {
	if time.Now().After(payload.RegisteredClaims.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}