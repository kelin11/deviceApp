package middleware

import (
	errmsg "deviceApp/code"
	"deviceApp/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var JwtKey = []byte(settings.JwtKey)
var str string

type Claims struct {
	Username string `json:"username"` //学号
	jwt.StandardClaims
}

func CreateToken(cliam Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliam)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, err
}

func JWTMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": errmsg.ERROR_TOKEN_WRONG, "msg": errmsg.GetMessage(errmsg.ERROR_TOKEN_WRONG)})
			ctx.Abort()
			return
		}
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			code := errmsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": errmsg.GetMessage(code)})
			ctx.Abort()
			return
		}
		ctx.Set("username", claims.Username)
	}
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	return token, Claims, err
}
