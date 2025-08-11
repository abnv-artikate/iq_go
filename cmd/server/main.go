package main

import (
	"log"
	"net/http"

	"iq-go/internal/auth"
	"iq-go/internal/config"
	"iq-go/internal/database"
	"iq-go/internal/handlers"
	"iq-go/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.RunMigrations(db)

	userService := services.NewUserService(db)
	testService := services.NewTestService(db)
	resultService := services.NewResultService(db)

	authHandler := handlers.NewAuthHandler(userService)
	testHandler := handlers.NewTestHandler(testService)
	resultHandler := handlers.NewResultHandler(resultService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r.LoadHTMLGlob("web/templates/*")
	r.Static("/static", "./web/static")

	// Web routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.GET("/test", auth.RequireAuth, func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", nil)
	})
	r.GET("/results", auth.RequireAuth, func(c *gin.Context) {
		c.HTML(http.StatusOK, "results.html", nil)
	})

	// API routes
	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.POST("/logout", authHandler.Logout)

		// Protected routes
		protected := api.Group("/")
		protected.Use(auth.RequireAuth)
		{
			protected.GET("/questions", testHandler.GetQuestions)
			protected.POST("/submit", testHandler.SubmitTest)
			protected.GET("/results", resultHandler.GetResults)
			protected.GET("/results/:id", resultHandler.GetResult)
		}
	}

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
