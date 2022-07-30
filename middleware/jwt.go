package middleware

import (
	errmsg "deviceApp/code"
	"deviceApp/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

//gin中设置控制jwt中间件拦截器 除了login都会进行登录拦截(进行token验证)
func JWTMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if strings.Contains(ctx.Request.URL.Path, "login") {
			//如果包含 那么久不进行拦截处理
			return
		}
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
		//中间设置全局变量username名称 在其他需要上报的时候直接在context中拿取即可
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
