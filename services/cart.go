package services

import (
	"context"
	"delivery-system/database"
	"delivery-system/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartService struct {
	Collection *mongo.Collection
}

func NewCartService() *CartService {
	collection := database.GetCollection("carts")
	return &CartService{Collection: collection}
}

func (cs *CartService) GetCartByUserID(userID primitive.ObjectID) (*models.Cart, error) {
	var cart models.Cart
	filter := bson.M{"user_id": userID}

	// Find the cart document
	err := cs.Collection.FindOne(context.TODO(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If no cart is found, return a new empty cart
			cart = models.Cart{UserID: userID, Items: []models.CartItem{}}
			return &cart, nil
		}
		return nil, err
	}

	return &cart, nil
}

func (cs *CartService) AddItemToCart(userID, productID primitive.ObjectID, quantity int) error {
	cart, err := cs.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity += quantity
			return cs.updateCart(cart)
		}
	}
	newItem := models.CartItem{
		ProductID: productID,
		Quantity:  quantity,
	}
	cart.Items = append(cart.Items, newItem)
	return cs.updateCart(cart)
}

func (cs *CartService) UpdateItemQuantity(userID, productID primitive.ObjectID, quantity int) error {
	cart, err := cs.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity = quantity
			return cs.updateCart(cart)
		}
	}
	return errors.New("product not found in cart")
}
func (cs *CartService) RemoveItemFromCart(userID, productID primitive.ObjectID) error {
	cart, err := cs.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	var updatedItems []models.CartItem
	var found bool
	for _, item := range cart.Items {
		if item.ProductID != productID {
			updatedItems = append(updatedItems, item)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("product not found in cart")
	}
	cart.Items = updatedItems
	return cs.updateCart(cart)
}

func (cs *CartService) ClearCart(userID primitive.ObjectID) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": bson.M{"items": []models.CartItem{}}}
	_, err := cs.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (cs *CartService) updateCart(cart *models.Cart) error {
	filter := bson.M{"user_id": cart.UserID}
	if len(cart.Items) < 1 {
		cart.Items = []models.CartItem{}
	}
	update := bson.M{"$set": bson.M{"items": cart.Items}}
	_, err := cs.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
