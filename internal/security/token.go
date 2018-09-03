package security

import (
	"fmt"
	"gitlab.com/dentych/dinner-dash/internal/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId int `json:"uid,omitempty"`
	jwt.StandardClaims
}

var ErrTokenExpired = fmt.Errorf("Token is expired")

func CreateJwtAccessToken(userId int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "dinnerdash-api",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Subject:   email,
		},
	})
	return token.SignedString([]byte("secret"))
}

func ValidateJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid token: OK: %v, Valid: %v", ok, token.Valid)
	}
}

func hmacSecret() jwt.Keyfunc {
	signingSecret := config.GetenvOrDefault("JWT_SECRET", "secret")

	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingSecret), nil
	}
}
