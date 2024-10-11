package categories

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	routesGroup := route.Group("/categories")
	// GET all categories
	routesGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all categories",
		})
	})

	// GET a specific category by ID
	routesGroup.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get category by ID",
		})
	})

	// POST a new category
	routesGroup.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create a new category",
		})
	})

	// PUT update an existing category
	routesGroup.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update a category",
		})
	})

	// DELETE a category
	routesGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete a category",
		})
	})
}
