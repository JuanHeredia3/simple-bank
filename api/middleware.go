package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/JuanHeredia3/simple-bank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "Bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorsResponse(fmt.Errorf("authorization header is not provided")))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorsResponse(fmt.Errorf("invalid authorization header format")))
			return
		}

		accessType := fields[0]
		if accessType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorsResponse(fmt.Errorf("unsupported authorization type")))
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorsResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
