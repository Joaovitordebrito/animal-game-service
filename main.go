package main

import (
	"animal-game/controllers"
	"animal-game/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/user", controllers.CreateUser)
	r.Run()
}
