package main

import (
	"project1/package/initializer"
	"project1/package/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_"project1/docs"
)

func init() {
	initializer.EnvLoad()
	initializer.LoadDatabase()
}

//	@title			E Commerce API
//	@version		1.0
//	@description	Ecommerce API in go using Gin frame work

//	@host		localhost:8080
//	@BasePath	/

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	user := router.Group("/")
	routers.UserGroup(user)

	admin := router.Group("/admin")
	routers.AdminGroup(admin)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

}
