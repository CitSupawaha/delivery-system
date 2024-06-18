package services

import (
	"context"
	"delivery-system/database"
	"delivery-system/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Collection *mongo.Collection
}

func NewProductService() *ProductService {
	collection := database.GetCollection("products")
	return &ProductService{Collection: collection}
}

func (ps *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := ps.Collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (ps *ProductService) GetProductsByCategoryID(categoryID primitive.ObjectID) ([]models.Product, error) {
	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"category_id": categoryID}
	cursor, err := ps.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) UpdateProductStock(productID primitive.ObjectID, newStock int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": productID}
	update := bson.M{"$set": bson.M{"stock": newStock}}

	_, err := ps.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
