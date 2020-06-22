package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gopress/models"
)

func InitAccount(router *gin.Engine) {
	router.GET("/login", Login)
	router.POST("/login", LoginPost)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context) {
	var user models.AccountLogin
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "welcome " + user.Username + " ,you are logged in"})
}
