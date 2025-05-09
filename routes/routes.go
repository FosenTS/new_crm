package routes

import (
	"crm/handlers"
	"crm/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handler) {
	// Public routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Product routes
		authorized.GET("/products", h.GetProducts)
		authorized.POST("/products", middleware.RoleMiddleware("boss"), h.CreateProduct)
		authorized.PUT("/products/:id", middleware.RoleMiddleware("boss"), h.UpdateProduct)

		// Order routes
		authorized.POST("/orders", h.CreateOrder)
		authorized.GET("/orders", h.GetOrders)
		authorized.PUT("/orders/:id/status", middleware.RoleMiddleware("boss"), h.UpdateOrderStatus)
	}
}
