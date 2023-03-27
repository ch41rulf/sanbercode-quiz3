package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"quiz-3/controllers"
)

func StartApp() *gin.Engine {

	users := map[string]string{
		"admin":  "password",
		"editor": "secret",
	}
	fmt.Println(users)

	router := gin.Default()
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.SegitigaSamaSisiHandler)
	router.GET("/bangun-datar/persegi", controllers.PersegiHandler)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjangHandler)
	router.GET("/bangun-datar/lingkaran", controllers.LingkaranHandler)

	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", controllers.InsertCategories)
	router.PUT("/categories/:id", controllers.UpdateCategories)
	router.DELETE("/categories/:id", controllers.DeleteCategories)
	router.GET("/categories/:id", controllers.GetBooksByCategoryIDHandler)

	router.GET("/books", controllers.GetBooks)
	router.Use(controllers.AuthMiddleware(users))
	router.POST("/books", controllers.InsertBooks)
	router.PUT("/books/:id", controllers.UpdateBooks) //GetBooks
	router.DELETE("/books/:id", controllers.DeleteBooks)

	return router
}
