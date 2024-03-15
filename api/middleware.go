package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dohaelsawy/bookStore/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		autherizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(autherizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponce(err))
			return
		}
		fields := strings.Fields(autherizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponce(err))
			return
		}

		autherizationType := strings.ToLower(fields[0])
		if autherizationType != authorizationTypeBearer {
			err := errors.New("authorization type is not matching")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponce(err))
			return
		}
		accessToken := fields[1]
		payload , err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponce(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
