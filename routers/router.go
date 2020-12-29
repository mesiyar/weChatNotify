package routers

import (
	//_ "cloud-notes/src/docs"
	//"cloud-notes/src/middleware/jwt"
	//v1 "cloud-notes/src/routers/api/v1"
	"github.com/gin-gonic/gin"
	"wechatNotify/middleware/jwt"
	"wechatNotify/pkg/setting"
	"wechatNotify/routers/api"
	v1 "wechatNotify/routers/api/v1"

	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	apiV1 := r.Group("/api/v1")

	apiV1.Use(jwt.JWT())
	{
		user := v1.UserController{}
		apiV1.GET("/users", user.GetUsers)
		apiV1.POST("/disable_account", api.DisableAuth)
		apiV1.POST("/create_account", api.CreateAuth)
		apiV1.POST("/logout", api.Logout)
		billing := v1.BillController{}
		apiV1.POST("/add_billing", billing.AddBilling)
	}

	return r
}
