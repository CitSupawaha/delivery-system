package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Items      []OrderItem        `json:"items,omitempty" bson:"items,omitempty"` // Change type to []primitive.ObjectID
	TotalPrice float64            `json:"total_price,omitempty" bson:"total_price,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
}

type Orders struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Items      []Product          `json:"items,omitempty" bson:"items,omitempty"` // Change type to []primitive.ObjectID
	TotalPrice float64            `json:"total_price,omitempty" bson:"total_price,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
}

type OrderItem struct {
	ProductID   primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	ImageURL    string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
}
