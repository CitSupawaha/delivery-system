// models/payment.go

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OrderID       primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	Amount        float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	Status        string             `json:"status,omitempty" bson:"status,omitempty"`
	PaymentMethod string             `json:"payment_method,omitempty" bson:"payment_method,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
