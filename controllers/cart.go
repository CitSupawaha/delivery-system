package controllers

import (
	"delivery-system/helper"
	"delivery-system/models"
	"delivery-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCart(c *gin.Context) {
	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	cartService := services.NewCartService()
	cart, err := cartService.GetCartByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}
	var newCart models.Carts
	newCart.ID = cart.ID
	newCart.UserID = cart.UserID
	productService := services.NewProductService()
	// Fetch product details for each item in the cart
	for _, item := range cart.Items {
		var cartItem models.CartProductItem
		product, _ := productService.GetProductByID(item.ProductID)
		cartItem.ID = product.ID
		cartItem.Quantity = item.Quantity
		cartItem.Description = product.Description
		cartItem.ImageURL = product.ImageURL
		cartItem.Name = product.Name
		cartItem.Price = product.Price
		cartItem.Stock = product.Stock
		newCart.Items = append(newCart.Items, cartItem)
	}
	c.JSON(http.StatusOK, helper.ResData(200, "success", "", newCart))
}

func AddItemToCart(c *gin.Context) {
	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	var input struct {
		ProductID string `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	productID, err := primitive.ObjectIDFromHex(input.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	cartService := services.NewCartService()
	err = cartService.AddItemToCart(user.ID, productID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusOK, helper.ResData(200, "success", "", input))
}

func RemoveItemFromCart(c *gin.Context) {
	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	var input struct {
		ProductID string `json:"product_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	productID, err := primitive.ObjectIDFromHex(input.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	cartService := services.NewCartService()
	err = cartService.RemoveItemFromCart(user.ID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, helper.ResData(200, "success", "", input))
}

func UpdateItemQuantity(c *gin.Context) {
	email, _ := c.Get("email")
	userService := services.NewUserService()
	user, err := userService.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	var input struct {
		ProductID string `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	productID, err := primitive.ObjectIDFromHex(input.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	cartService := services.NewCartService()
	err = cartService.UpdateItemQuantity(user.ID, productID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item quantity"})
		return
	}

	c.JSON(http.StatusOK, helper.ResData(200, "success", "", input))
}
