// services/payment_service.go

package services

import (
	"context"
	"delivery-system/database"
	"delivery-system/models"
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentService struct {
	Collection *mongo.Collection
}

func NewPaymentService() *PaymentService {
	collection := database.GetCollection("payments")
	return &PaymentService{Collection: collection}
}

func (ps *PaymentService) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	payment.ID = primitive.NewObjectID()
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = payment.CreatedAt

	_, err := ps.Collection.InsertOne(context.TODO(), payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (ps *PaymentService) GetPaymentByID(paymentID primitive.ObjectID) (*models.Payment, error) {
	var payment models.Payment
	filter := bson.M{"_id": paymentID}
	err := ps.Collection.FindOne(context.TODO(), filter).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (ps *PaymentService) CreatePaymentIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
	// Ensure the amount is above the minimum threshold for the currency
	if amount < 1000 && currency == "THB" {
		return nil, fmt.Errorf("amount too small: Amount must be at least 1000 satang for THB")
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}
	return pi, nil
}

// services/payment_service.go

// func (ps *PaymentService) ProcessAndCreatePayment(payment *models.Payment) (*models.Payment, error) {
//     paymentGateway := NewPaymentGateway()
//     transactionID, err := paymentGateway.ProcessPayment(payment)
//     if err != nil {
//         return nil, err
//     }

//     payment.Status = "Completed"
//     payment.TransactionID = transactionID

//     return ps.CreatePayment(payment)
// }
