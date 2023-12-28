package middleware

import (
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
		}

		userID, err := getUserId(userToken)
		if err != nil {
			//if !errors.Is(err, jwt.ErrTokenNotValidYet) {
			//	ctx.AbortWithStatus(http.StatusUnauthorized)
			//}

			errToken := setToken(ctx)
			if errToken != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}

		ctx.Set("userID", userID)
	}
}

func setToken(ctx *gin.Context) error {
	token, err := CreateToken()
	if err != nil {
		return err
	}

	ctx.SetCookie("UserToken", token, int(TOKEN_EXP.Seconds()), "/", "localhost", false, true)
	return nil
}

func CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		UserID: uuid.New().String(),
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", fmt.Errorf("token signing :%w", err)
	}

	return tokenString, nil
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
