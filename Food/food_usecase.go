package Food

import (
	entity "GladiResik/Entity"
	"context"
)

type FoodUsecase struct {
	Datalayer *FoodData
}

func (FUC *FoodUsecase) CheckRange(ctx context.Context, id int) bool {
	min, max := FUC.Datalayer.GetRange(ctx)
	if id < min || id > max {
		return false
	}
	return true
}

func FUC_Init(ctx context.Context, data *FoodData) *FoodUsecase {
	return &FoodUsecase{
		Datalayer: data,
	}
}

func (FUC *FoodUsecase) ViewAll(ctx context.Context) ([]entity.Food, error) {
	return FUC.Datalayer.ViewAll(ctx)
}

func (FUC *FoodUsecase) ViewOne(ctx context.Context, id int) (entity.Food, error) {
	var food entity.Food
	if !FUC.CheckRange(ctx, id) {
		return food, ctx.Err()
	}
	return FUC.Datalayer.ViewOne(ctx, id)
}

func (FUC *FoodUsecase) UpdateStatus(ctx context.Context, id int, status entity.Food) error {
	if !FUC.CheckRange(ctx, id) {
		return ctx.Err()
	}
	food, _ := FUC.Datalayer.ViewOne(ctx, id)
	food.Status = status.Status
	return FUC.Datalayer.Update(ctx, id, food)
}
