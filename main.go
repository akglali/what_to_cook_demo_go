package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"what_to_cook_demo_go/database"
	"what_to_cook_demo_go/user"
)

func main() {
	//call the database connection at the beginning
	database.ConnectDatabase()
	// don't want to get warning every time I run
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	userGroup := router.Group("/User")
	user.SetupEntry(userGroup)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
}
