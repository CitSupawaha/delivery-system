package models

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	CategoryID  primitive.ObjectID `json:"category_id,omitempty" bson:"category_id,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
	ImageURL    string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
}

type CreateProduct struct {
	ID          primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string                `form:"name,omitempty" bson:"name,omitempty"`
	Description string                `form:"description,omitempty" bson:"description,omitempty"`
	Price       float64               `form:"price,omitempty" bson:"price,omitempty"`
	CategoryID  primitive.ObjectID    `form:"category_id,omitempty" bson:"category_id,omitempty"`
	Stock       int                   `form:"stock,omitempty" bson:"stock,omitempty"`
	ImageURL    *multipart.FileHeader `form:"image_url,omitempty" bson:"image_url,omitempty"`
}
