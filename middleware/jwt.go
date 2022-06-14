package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//中间件


var JwtKey = []byte(utils.JwtKey)
var code int

//MyClaims 我的Claims 里面要包含jwt.StandardClaims结构
type MyClaims struct {
	Username string		`json:"username"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string)(string,int){
	//过期时间，当前时间后的10小时为过期时间，用jwt.at函数转为*Time
	expireTime := jwt.At(time.Now().Add(10 * time.Hour))
	//填充MyClaims
	SetClaims := MyClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,    //过期时间
			Issuer: "ginblog",		  //签发人
		},
	}

	//根据MyClaims，生成token； SigningMethodHS256为加密方法
	reqToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	tokenStr,err := reqToken.SignedString(JwtKey)  //返回完整的字符串token
	if err != nil {
		return "",errmsg.ERROR
	}
	return tokenStr,errmsg.SUCCESS
}

//CheckToken 验证token
func CheckToken(token string)(*MyClaims,int){
	//解析token字符串，拿到token
	setToken,_ := jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
								//类型断言
	if key,_ := setToken.Claims.(*MyClaims); setToken.Valid{
		return key,errmsg.SUCCESS
	}else{
		return nil,errmsg.ERROR
	}
}
//JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		//如果没有与Authorization关联的值，则返回空
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder," ",2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key,tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR{
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt.Unix(){
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}


		c.Set("username",key.Username)
		c.Next()
	}
}
