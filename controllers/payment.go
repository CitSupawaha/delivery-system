// controllers/payment_controller.go

package controllers

import (
	"delivery-system/models"
	"delivery-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
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

	payment.UserID = user.ID
	paymentService := services.NewPaymentService()
	createdPayment, err := paymentService.CreatePayment(&payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": createdPayment})
}

func GetPayment(c *gin.Context) {
	paymentID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(paymentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	paymentService := services.NewPaymentService()
	payment, err := paymentService.GetPaymentByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

// controllers/payment_controller.go

func CreatePaymentIntent(c *gin.Context) {
	var req struct {
		Amount   int64  `json:"amount" binding:"required"`
		Currency string `json:"currency" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	paymentService := services.NewPaymentService()
	paymentIntent, err := paymentService.CreatePaymentIntent(req.Amount, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_secret": paymentIntent.ClientSecret})
}
