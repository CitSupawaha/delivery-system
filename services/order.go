package services

import (
	"context"
	"delivery-system/database"
	"delivery-system/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	Collection *mongo.Collection
}

func NewOrderService() *OrderService {
	collection := database.GetCollection("orders")
	return &OrderService{Collection: collection}
}

func (os *OrderService) CreateOrder(order *models.Order) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := os.Collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID = result.InsertedID.(primitive.ObjectID)
	return order, nil
}

func (os *OrderService) GetOrdersByUserID(userID primitive.ObjectID) ([]models.Orders, error) {
	var orders []models.Orders

	ctx := context.TODO()
	filter := bson.M{"user_id": userID}

	cursor, err := os.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order models.Order
		var newOrder models.Orders
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		var orderID []primitive.ObjectID
		for _, item := range order.Items {
			orderID = append(orderID, item.ProductID)
		}
		// Fetch detailed item information for the order
		items, err := os.fetchOrderItems(ctx, orderID)
		if err != nil {
			return nil, err
		}
		newOrder.Items = items
		newOrder.ID = order.ID
		newOrder.Status = order.Status
		newOrder.TotalPrice = order.TotalPrice
		newOrder.UserID = order.UserID
		orders = append(orders, newOrder)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (os *OrderService) fetchOrderItems(ctx context.Context, itemIDs []primitive.ObjectID) ([]models.Product, error) {
	var items []models.Product

	// Example: Fetch items from the products collection based on product IDs
	productsCollection := database.GetCollection("products")

	filter := bson.M{"_id": bson.M{"$in": itemIDs}}

	cursor, err := productsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		items = append(items, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (ps *ProductService) GetProductByID(productID primitive.ObjectID) (*models.Product, error) {
	ctx := context.TODO()
	var product models.Product
	filter := bson.M{"_id": productID}
	err := ps.Collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *ProductService) DecrementStock(productID primitive.ObjectID, quantity int) error {
	ctx := context.TODO()
	filter := bson.M{"_id": productID}
	update := bson.M{"$inc": bson.M{"stock": -quantity}}

	result, err := ps.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (os *OrderService) CreateOrderFromCart(userID primitive.ObjectID) (*models.Order, error) {
	cartService := NewCartService()
	cart, err := cartService.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	productService := NewProductService()
	var orderItems []models.OrderItem
	var totalPrice float64

	for _, cartItem := range cart.Items {
		product, err := productService.GetProductByID(cartItem.ProductID)
		if err != nil {
			return nil, err
		}
		if product.Stock < cartItem.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		orderItem := models.OrderItem{
			ProductID:   product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    cartItem.Quantity,
			ImageURL:    product.ImageURL,
		}
		orderItems = append(orderItems, orderItem)
		totalPrice += product.Price * float64(cartItem.Quantity)
	}

	order := models.Order{
		UserID:     userID,
		Items:      orderItems,
		TotalPrice: totalPrice,
		Status:     "Pending",
	}

	result, err := os.Collection.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}
	order.ID = result.InsertedID.(primitive.ObjectID)

	for _, cartItem := range cart.Items {
		err := productService.DecrementStock(cartItem.ProductID, cartItem.Quantity)
		if err != nil {
			return nil, err
		}
	}

	err = cartService.ClearCart(userID)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
