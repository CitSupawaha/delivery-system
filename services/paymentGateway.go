// services/payment_gateway.go

package services

import (
	"delivery-system/models"
	"errors"
)

type PaymentGateway struct{}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (pg *PaymentGateway) ProcessPayment(payment *models.Payment) (string, error) {
	// Simulate payment processing
	if payment.Amount <= 0 {
		return "", errors.New("invalid payment amount")
	}

	// Here you would call the real payment gateway API
	transactionID := "fake-transaction-id" // This should be returned by the payment gateway

	return transactionID, nil
}
