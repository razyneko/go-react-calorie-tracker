package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/razyneko/go-react-calorie-tracker/routes"
)

func main() {
	//setting up the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// create router
	router := gin.New()
	// middleware , when which an api is called or if it gives error
	//Use attaches global middleware to router
	router.Use(gin.Logger())
	router.Use(cors.Default())

	//routes
	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)
	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	//run the server at port
	router.Run(":" + port)
}
