package Orders

import (
	entity "GladiResik/Entity"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderData struct {
	Database *mongo.Collection
}

func ODL_Init(ctx context.Context, DB *mongo.Database) *OrderData {
	return &OrderData{
		Database: DB.Collection("Orders"),
	}
}

func (Orders *OrderData) Insert(ctx context.Context, item entity.Orders) {
	_, err := Orders.Database.InsertOne(ctx, item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}

func (Orders *OrderData) ViewAll(ctx context.Context) ([]entity.Orders, error) {
	var mainOrder []entity.Orders

	filter := bson.D{}
	cursor, err := Orders.Database.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order entity.Orders
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		mainOrder = append(mainOrder, order)
	}

	return mainOrder, nil
}
