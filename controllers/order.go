package controllers

import (
	"delivery-system/helper"
	"delivery-system/models"
	"delivery-system/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	order.UserID = user.ID

	var orderItems []models.OrderItem
	for _, item := range order.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	order.Items = orderItems

	orderService := services.NewOrderService()
	createdOrder, err := orderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	productService := services.NewProductService()
	for _, item := range order.Items {
		err := productService.DecrementStock(item.ProductID, item.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"order": createdOrder})
}

func GetOrdersByUser(c *gin.Context) {
	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	orderService := services.NewOrderService()
	orders, err := orderService.GetOrdersByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	fmt.Println("order ==> ", orders)

	productService := services.NewProductService()
	for i, order := range orders {
		for j, item := range order.Items {
			fmt.Println("item ==> ", item)
			product, err := productService.GetProductByID(item.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product details"})
				return
			}
			fmt.Println("proded=uct ==> ", product)
			orders[i].Items[j] = *product
		}
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func CreateOrderFromCart(c *gin.Context) {
	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	orderService := services.NewOrderService()
	order, err := orderService.CreateOrderFromCart(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helper.ResData(200, "success", "", order))
}
