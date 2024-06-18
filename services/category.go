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

type CategoryService struct {
	Collection *mongo.Collection
}

func NewCategoryService() *CategoryService {
	collection := database.GetCollection("categories")
	return &CategoryService{Collection: collection}
}

func (cs *CategoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := cs.Collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)
	return category, nil
}

func (cs *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := cs.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category models.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
