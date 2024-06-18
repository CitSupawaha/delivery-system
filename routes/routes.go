package routes

import (
	"delivery-system/controllers"
	"delivery-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/users", controllers.CreateUser)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// authorized.POST("/orders", controllers.CreateOrder)
		authorized.GET("/users", controllers.GetAllUsers)

		authorized.POST("/products", controllers.CreateProduct)

		authorized.POST("/category", controllers.CreateCategory)
		authorized.GET("/category", controllers.GetAllCategories)

		//authorized.POST("/orders", controllers.CreateOrder)
		authorized.POST("/orders", controllers.CreateOrderFromCart)
		authorized.GET("/orders", controllers.GetOrdersByUser)
		authorized.GET("/cart", controllers.GetCart)

		authorized.POST("/cart/add", controllers.AddItemToCart)
		authorized.POST("/cart/update", controllers.UpdateItemQuantity)
		authorized.POST("/cart/remove", controllers.RemoveItemFromCart)

		authorized.POST("/payments", controllers.CreatePayment)
		authorized.GET("/payments/:id", controllers.GetPayment)
		authorized.POST("/payments/intent", controllers.CreatePaymentIntent)
	}
}
