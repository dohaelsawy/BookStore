package token

import "time"

type Maker interface {
	// this function takes a username and the time duration and return a token and error
	CreateToken(username string, duration time.Duration) ( string, error)
	VerifyToken(token string) (*Payload,error)
}