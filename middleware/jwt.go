package middleware

import (
	"gineasy/conf"
	//jwt "OnlinePhotoAlbum/test"
	"gineasy/views"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"

	"time"
)
//init user's token
func GetToken() *jwt.GinJWTMiddleware {
	UserToken, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         conf.JwtRealm,
		Key:           []byte(conf.JwtKey),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "id",
		Unauthorized:  unauthorized,
	})
	if err != nil {
		log.Fatal("JWT ErrorBind:" + err.Error())
		return nil
	} else {
		return UserToken
	}
}

//return login failed message
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(200, views.ErrorRes("jwt鉴权错误："+message+"，请尝试重新登录"))
}
