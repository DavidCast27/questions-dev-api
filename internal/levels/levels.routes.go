package levels

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	routesGroup := route.Group("/levels")

	// GET all levels
	routesGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all levels",
		})
	})

	// GET a specific level by ID
	routesGroup.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get level by ID",
		})
	})

	// POST a new level
	routesGroup.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create a new level",
		})
	})

	// PUT update an existing level
	routesGroup.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update a level",
		})
	})

	// DELETE a level
	routesGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete a level",
		})
	})
}
