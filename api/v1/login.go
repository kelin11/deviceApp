package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/middleware"
	"deviceApp/model"
	"deviceApp/model/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Username string

//维修人员的账号数据库里面有吗?
func Login(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user)
	code, username := orm.CheckLogin(&user)
	if code == errmsg.ERROR {
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetMessage(code),
		})
		return
	}
	Username = username
	SetToken(ctx, user)
}

func SetToken(c *gin.Context, user model.User) {
	claims := middleware.Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "deviceApp",
		},
	}
	token, err := middleware.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_TOKEN_WRONG,
			"message": errmsg.GetMessage(errmsg.ERROR_TOKEN_WRONG),
			"token":   token,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,K
		"data":    user.UserName,
		"message": errmsg.GetMessage(200),
		"token":   token,
		//"id":      user.ID,
	})
	return
}
