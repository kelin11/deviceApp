package middleware

import (
	v1 "deviceApp/api/v1"
	errmsg "deviceApp/code"
	"deviceApp/settings"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var JwtKey = []byte(settings.JwtKey)
var str string

type Claims struct {
	Username string `json:"username"` //学号
	jwt.StandardClaims
}

func CreateToken(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	cliams := &Claims{
		v1.Username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1", // 签名颁发者
			Subject:   "deviceApp", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
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
		v1.Username = claims.Username
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
