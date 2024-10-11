package main

import (
	"github.com/DavidCast27/questions-dev-api/internal/answers"
	"github.com/DavidCast27/questions-dev-api/internal/categories"
	"github.com/DavidCast27/questions-dev-api/internal/languages"
	"github.com/DavidCast27/questions-dev-api/internal/levels"
	"github.com/DavidCast27/questions-dev-api/internal/questions"
	"github.com/DavidCast27/questions-dev-api/internal/scores"
	"github.com/DavidCast27/questions-dev-api/internal/tags"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	answers.Routes(route)
	categories.Routes(route)
	languages.Routes(route)
	levels.Routes(route)
	questions.Routes(route)
	scores.Routes(route)
	tags.Routes(route)

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.Run()

}
