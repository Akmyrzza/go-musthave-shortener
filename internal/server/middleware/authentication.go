package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

const TOKEN_EXP = time.Hour
const SECRET_KEY = "supersecretkey"

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userToken, err := ctx.Cookie("UserToken")
		if err != nil {
			errToken := setToken(ctx)
			if errToken != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			return
		}

		userID, err := getUserId(userToken)
		if err != nil {
			if !errors.Is(err, jwt.ErrTokenNotValidYet) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			errToken := setToken(ctx)
			if errToken != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}

		ctx.Set("userID", userID)
	}
}

func setToken(ctx *gin.Context) error {
	token, userID, err := CreateToken()
	if err != nil {
		return err
	}

	ctx.SetCookie("UserToken", token, int(TOKEN_EXP.Seconds()), "/", "localhost", false, true)
	ctx.Set("userID", userID)
	ctx.Set("newUser", true)
	return nil
}

func CreateToken() (string, string, error) {
	userID := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", fmt.Errorf("token signing :%w", err)
	}

	return tokenString, userID, nil
}

func getUserId(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})
	if err != nil {
		return "", fmt.Errorf("token parsing: %w", err)
	}

	if !token.Valid {
		return "", jwt.ErrTokenNotValidYet
	}

	return claims.UserID, nil
}
