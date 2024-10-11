package tags

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	routesGroup := route.Group("/tags")

	// GET all tags
	routesGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all tags",
		})
	})

	// GET a specific tag by ID
	routesGroup.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get tag by ID",
		})
	})

	// POST a new tag
	routesGroup.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create a new tag",
		})
	})

	// PUT update an existing tag
	routesGroup.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update a tag",
		})
	})

	// DELETE a tag
	routesGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete a tag",
		})
	})
}
