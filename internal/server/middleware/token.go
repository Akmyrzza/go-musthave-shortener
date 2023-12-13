package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type AuthorizationToken struct {
	UserID  string
	Expires time.Time
}

var Secret string = "mysecret"

func TokenCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie, err := ctx.Cookie("uuid")
		if err != nil {
			newTokenCookie(ctx)
			return
		}

		_, err = uuid.Parse(cookie)
		if err != nil {
			newTokenCookie(ctx)
			return
		}
	}
}

func newTokenCookie(ctx *gin.Context) {
	userID := uuid.New().String()

	h := hmac.New(sha256.New, []byte(Secret))
	h.Write([]byte(userID))

	token := hex.EncodeToString(h.Sum(nil))

	ctx.SetCookie("user_id", userID, 3600, "/", "localhost", false, true)
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
	ctx.AbortWithStatus(http.StatusUnauthorized)
}
