package main

import (
	"awesomeProject2/database"
	"awesomeProject2/requestHandler"
	"awesomeProject2/settings"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	options := settings.Load("settings/serverSetting.json")
	e := database.Connect(options)
	if e != nil {
		fmt.Println(e)
		return
	}

	user := database.User{}
	users := user.SelectAll()

	fmt.Println(users)

	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.Static("assets", "assets")
	router.GET("/", requestHandler.Index)
	_ = router.Run(options.Address + ":" + options.Port)
}
