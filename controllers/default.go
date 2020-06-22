package controllers

import "github.com/gin-gonic/gin"

func InitController(router *gin.Engine) {
	InitAccount(router)
}
