package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cart represents a user's shopping cart
type Cart struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Items  []CartItem         `json:"items,omitempty" bson:"items,omitempty"`
}

type Carts struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Items  []CartProductItem  `json:"items,omitempty" bson:"items,omitempty"`
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	ProductID primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity  int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

// CartItem represents an item in the shopping cart
type CartProductItem struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	CategoryID  primitive.ObjectID `json:"category_id,omitempty" bson:"category_id,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
	ImageURL    string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
