package main

import (
	"LoginAndRegisterApiJWT/controllers"
	"LoginAndRegisterApiJWT/database"
	"LoginAndRegisterApiJWT/middlewares"
	"LoginAndRegisterApiJWT/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()

	r.Run()
}

// setupRouter sets up the router and adds the routes.
func setupRouter() *gin.Engine {
	// Create a new router
	r := gin.Default()
	// Add a welcome route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	// Create a new group for the API
	api := r.Group("/api")
	{
		// Create a new group for the public routes
		public := api.Group("/public")
		{
			// Add the login route
			public.POST("/login", controllers.Login)
			// Add the signup route
			public.POST("/signup", controllers.Signup)
		}
		// Add the signup route
		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			// Add the profile route
			protected.GET("/profile", controllers.Profile)
		}
	}
	// Return the router
	return r
}
