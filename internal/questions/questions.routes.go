package questions

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	routesGroup := route.Group("/questions")

	// GET all questions
	routesGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all questions",
		})
	})

	// GET a specific question by ID
	routesGroup.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get question by ID",
		})
	})

	// POST a new question
	routesGroup.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create a new question",
		})
	})

	// PUT update an existing question
	routesGroup.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update a question",
		})
	})

	// DELETE a question
	routesGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete a question",
		})
	})

	//  ------------  TAGS ------------

	// GET all tags related to a question
	routesGroup.GET(":id/tags", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all tags ",
		})
	})

	//DELETE Disassociate a tag from a question.
	routesGroup.DELETE(":id/tags/:tag", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "remove tag from a question",
		})
	})

	//POST add a tag with a question
	routesGroup.POST(":id/tags/:tag", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "join tag with a question",
		})
	})

}
