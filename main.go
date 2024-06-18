package main

import (
	"delivery-system/database"
	"delivery-system/routes"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
)

func main() {
	router := gin.Default()
	database.Connect()
	stripe.Key = "sk_test_51PSw8V03uQKPSfOrEAEuD6mKCbGILdh0agrmKYWZu8u0TbaLAdrAU7rNkefvgoUCW7Ye7K1yNW9ELmJ7DxUzACLW00pDLgLDoS"
	routes.SetupRoutes(router)
	router.Run(":8000")
}
