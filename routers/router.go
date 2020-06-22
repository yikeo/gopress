package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gopress/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	controllers.InitController(router)
	return router
}
