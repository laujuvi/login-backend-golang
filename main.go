package main

import (
	"github.com/laujuvi/login-system/internal/auth"
	"github.com/laujuvi/login-system/internal/database"
	"github.com/laujuvi/login-system/internal/middleware"
	"github.com/laujuvi/login-system/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	auth.Init()

	r := gin.Default()

	// Rutas publicas
	r.POST("/register", func(c *gin.Context) {
		user.Register(c, database.DB)
	})

	r.POST("/login", auth.LoginHandler)

	r.POST("/auth/refresh", auth.Refresh)

	// Rutas protegidas
	protected := r.Group("/auth")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			email := c.GetString("email")
			userID := c.GetString("user_id")

			c.JSON(200, gin.H{
				"message": "Token válido",
				"user_id": userID,
				"email":   email,
			})
		})

	}

	r.Run(":8000")

}
