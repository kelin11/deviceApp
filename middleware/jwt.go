package middleware

import (
	"deviceApp/settings"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{[]byte(settings.JwtKey)}
}

func (j *JWT) CreateJwt(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(settings.JwtKey)
}

func (j *JWT) ParseJwt(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	return nil, err
}
