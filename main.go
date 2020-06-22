package main

import (
	"github.com/gopress/routers"
)

func main() {

	// gin.DisableConsoleColor()

	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	router := routers.InitRouter()
	router.Static("/static", "./static")
	router.Run(":80")
}
