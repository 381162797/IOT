package router

import (
	"mserver/router/api"
	v1 "mserver/router/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Init 返回router 并开启WebSocket服务
func Init() (r *gin.Engine) {
	r = gin.Default()
	r.Use(cors.Default())
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	rAPI := r.Group("api")
	{
		rAPI.GET("auth", api.GetAuth)
		rAPI.POST("user", v1.UserRegister)
		rAPI.DELETE("user", v1.DeleteUser)
		rAPI.PUT("user", v1.PutUser)
		rAPI.POST("login", v1.Login)
	}
	// 使用 JWT
	// rAPI.Use(jwt.AuthJWT())
	apiv1 := rAPI.Group("v1")
	{
		//device
		apiv1.GET("/device", v1.GetDevice)
		apiv1.POST("/device", v1.AddDevice)
		apiv1.PUT("/device", v1.UpdateDevice)
		apiv1.DELETE("/device", v1.DeleteDevice)
		//template
		apiv1.GET("/template", v1.GetTemplate)
		apiv1.POST("/template", v1.AddTemplate)
		apiv1.PUT("/template", v1.UpdateTemplate)
		apiv1.DELETE("/template", v1.DeleteTemplate)
		//schedule
		apiv1.GET("/schedule", v1.GetSchedule)
		apiv1.POST("/schedule", v1.AddSchedule)
		apiv1.PUT("/schedule", v1.UpdateSchedule)
		apiv1.DELETE("/schedule", v1.DeleteSchedule)
		//pointdata
		apiv1.GET("/data", v1.GetPointData)
		apiv1.DELETE("/data", v1.DeletePointData)
	}
	return
}
