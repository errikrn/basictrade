package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

var secretKey = "your-256-bit-secret"

func GenereteToken(id uint, email string) string {
	claims := jwt5.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 100),
	}

	parseToken := jwt5.NewWithClaims(jwt5.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return err.Error()
	}

	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt5.Parse(stringToken, func(t *jwt5.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt5.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	claims, ok := token.Claims.(jwt5.MapClaims)
	if !ok && !token.Valid {
		return nil, errResponse
	}

	expClaim, exists := claims["exp"]
	if !exists {
		return nil, errors.New("expire claim is missing")
	}

	expStr, ok := expClaim.(string)
	if !ok {
		return nil, errors.New("expire claim is not a valid type")
	}

	expTime, err := time.Parse(time.RFC3339, expStr)
	if err != nil {
		return nil, errors.New("error parsing expiration time")
	}

	if time.Now().After(expTime) {
		return nil, errors.New("token is expired")
	}

	return token.Claims.(jwt5.MapClaims), nil
}
