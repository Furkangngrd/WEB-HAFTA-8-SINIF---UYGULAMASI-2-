package main

import (
	"log"

	"golearn/config"
	"golearn/database"
	"golearn/handlers"
	"golearn/middleware"
	"golearn/websocket"

	"github.com/gin-gonic/gin"

	_ "golearn/docs/swagger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoLearn API
// @version 1.0
// @description Uzaktan Eğitim Platformu API
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg.DBPath)

	r := gin.Default()

	r.Use(middleware.RateLimiter())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register(db))
			auth.POST("/login", handlers.Login(db, cfg.JWTSecret))
		}

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			// Courses
			courses := protected.Group("/courses")
			{
				courses.GET("", handlers.GetCourses(db))
				courses.GET("/:id", handlers.GetCourse(db))
				courses.POST("", middleware.RoleMiddleware("teacher"), handlers.CreateCourse(db))
				courses.PUT("/:id", middleware.RoleMiddleware("teacher"), handlers.UpdateCourse(db))
				courses.DELETE("/:id", middleware.RoleMiddleware("teacher"), handlers.DeleteCourse(db))
				courses.POST("/:id/enroll", middleware.RoleMiddleware("student"), handlers.EnrollCourse(db))
			}

			// Lessons
			lessons := protected.Group("/courses/:id/lessons")
			{
				lessons.GET("", handlers.GetLessons(db))
				lessons.GET("/:lessonId", handlers.GetLesson(db))
				lessons.POST("", middleware.RoleMiddleware("teacher"), handlers.CreateLesson(db))
				lessons.PUT("/:lessonId", middleware.RoleMiddleware("teacher"), handlers.UpdateLesson(db))
				lessons.DELETE("/:lessonId", middleware.RoleMiddleware("teacher"), handlers.DeleteLesson(db))
			}

			// Quizzes
			quizzes := protected.Group("/courses/:id/quizzes")
			{
				quizzes.GET("", handlers.GetQuizzes(db))
				quizzes.GET("/:quizId", handlers.GetQuiz(db))
				quizzes.POST("", middleware.RoleMiddleware("teacher"), handlers.CreateQuiz(db))
				quizzes.POST("/:quizId/submit", middleware.RoleMiddleware("student"), handlers.SubmitQuiz(db))
				quizzes.GET("/:quizId/results", handlers.GetQuizResults(db))
			}

			// Progress
			progress := protected.Group("/progress")
			{
				progress.POST("/complete", middleware.RoleMiddleware("student"), handlers.CompleteLesson(db))
				progress.GET("/courses/:id", handlers.GetProgress(db))
			}
		}

		// WebSocket
		hub := websocket.NewHub()
		go hub.Run()
		api.GET("/ws/:roomId", middleware.AuthMiddleware(cfg.JWTSecret), handlers.HandleWebSocket(hub))
	}

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
