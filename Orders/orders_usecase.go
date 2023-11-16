package Orders

import (
	entity "GladiResik/Entity"
	"GladiResik/Food"
	"context"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return newUUID.String()
}

type OrderUsecase struct {
	Datalayer *OrderData
	FoodData  *Food.FoodData
}

func OUC_Init(ctx context.Context, data *OrderData, food *Food.FoodData) *OrderUsecase {
	return &OrderUsecase{
		FoodData:  food,
		Datalayer: data,
	}
}

func (usecase *OrderUsecase) Create(ctx context.Context, order []entity.Incoming) {
	var NewOrder entity.Orders
	for _, x := range order {
		NewOrder.Uuid = GenerateUUID()
		food, _ := usecase.FoodData.ViewOne(ctx, x.FoodID)
		NewOrder.Contents = append(NewOrder.Contents, food)
		NewOrder.Quantity = append(NewOrder.Quantity, x.Quantity)
		NewOrder.Completion = append(NewOrder.Completion, false)

	}
	usecase.Datalayer.Insert(ctx, NewOrder)
}

func (usecase *OrderUsecase) ViewAll(ctx context.Context) []entity.Orders {
	item, _ := usecase.Datalayer.ViewAll(ctx)
	return item
}
