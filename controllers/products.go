package controllers

import (
	"delivery-system/models"
	"delivery-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	productService := services.NewProductService()

	// Create product in database
	createdProduct, err := productService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": createdProduct})
}

func UpdateProductStock(c *gin.Context) {
	var updateStockRequest struct {
		ProductID string `json:"product_id"`
		NewStock  int    `json:"new_stock"`
	}

	if err := c.ShouldBindJSON(&updateStockRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	productID, err := primitive.ObjectIDFromHex(updateStockRequest.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	productService := services.NewProductService()
	if err := productService.UpdateProductStock(productID, updateStockRequest.NewStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product stock updated successfully"})
}
