package languages

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	routesGroup := route.Group("/languages")

	// GET all languages
	routesGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all languages",
		})
	})

	// GET a specific language by ID
	routesGroup.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get language by ID",
		})
	})

	// POST a new language
	routesGroup.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create a new language",
		})
	})

	// PUT update an existing language
	routesGroup.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update a language",
		})
	})

	// DELETE a language
	routesGroup.DELETE("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete a language",
		})
	})
}
