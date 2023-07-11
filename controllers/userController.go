package controllers

import (
	"animal-game/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct {
		UserName string
	}

	c.Bind(&body)
	fmt.Println(body)
	err := service.CreateUser(body.UserName)

	if err != nil {
		c.Status(400)
		return
	}
}
