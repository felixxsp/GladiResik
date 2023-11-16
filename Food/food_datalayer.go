package Food

import (
	entity "GladiResik/Entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodData struct {
	foods mongo.Collection
}

func FDL_Init(ctx context.Context, database *mongo.Database) *FoodData {
	return &FoodData{
		foods: *database.Collection("Food"),
	}
}

func (repo *FoodData) GetRange(ctx context.Context) (int, int) {
	foods, _ := repo.ViewAll(ctx)
	min := 500
	max := 0
	for _, x := range foods {
		if x.ID > max {
			max = x.ID
		}
		if x.ID < min {
			min = x.ID
		}
	}
	return min, max
}

func (repo *FoodData) ViewAll(ctx context.Context) ([]entity.Food, error) {
	var foods []entity.Food

	filter := bson.D{}
	cursor, err := repo.foods.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var food entity.Food
		err := cursor.Decode(&food)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	return foods, nil
}

func (repo *FoodData) ViewOne(ctx context.Context, id int) (entity.Food, error) {
	var food entity.Food

	cursor, err := repo.foods.Find(ctx, bson.M{"id": 1})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	cursor.Decode(&food)
	return food, nil
}

func (repo *FoodData) Update(ctx context.Context, id int, status entity.Food) error {
	_, err := repo.foods.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": status},
	)
	return err
}
