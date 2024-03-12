package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	//uuid "github.com/google/uuid"
)

const (
	minSecretKey = 32
)

type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKey {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKey)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error in type assertion for claims")
	}
	id , ok := claims["jti"].(string)
	if !ok {
		return nil, fmt.Errorf("%v , %v", id , err)
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("error in extracting username from claims")
	}
	
	expiresAt , err := claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("expires not found in claims")
	}
	issuedAt , err := claims.GetIssuedAt()
	if err != nil {
		return nil, fmt.Errorf("issued not found in claims")
	}

	// Add any other payload fields here

	// Create a new Payload instance
	payload := &Payload{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt:expiresAt,
			IssuedAt:  issuedAt,
			ID: id,
		},
	}

	return payload, nil
}
