package controllers

import (
	"delivery-system/helper"
	"delivery-system/models"
	"delivery-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	categoryService := services.NewCategoryService()
	createdCategory, err := categoryService.CreateCategory(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, helper.ResData(200, "success", "", createdCategory))
}

func GetAllCategories(c *gin.Context) {
	categoryService := services.NewCategoryService()
	categories, err := categoryService.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, helper.ResData(200, "success", "", categories))
}
