package main

import (
	food "GladiResik/Food"
	"GladiResik/Orders"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mainCtx := context.Background()

	router := gin.Default()

	client, _ := mongo.Connect(mainCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	database := client.Database("GladiResik")

	FoodData := food.FDL_Init(mainCtx, database)
	FUC := food.FUC_Init(mainCtx, FoodData)
	FoodHandler := food.FHD_Init(router, FUC)
	FoodHandler.Standby(mainCtx)

	OrderData := Orders.ODL_Init(mainCtx, database)
	OrderUC := Orders.OUC_Init(mainCtx, OrderData, FoodData)
	OrderHandler := Orders.OHD_Init(router, OrderUC)
	OrderHandler.Standby(mainCtx)

	router.Run(":8080")
}
