package Food

import (
	entity "GladiResik/Entity"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodHandlers struct {
	Router *gin.Engine
	FUC    *FoodUsecase
}

func FHD_Init(router *gin.Engine, FUC *FoodUsecase) *FoodHandlers {
	return &FoodHandlers{
		Router: router,
		FUC:    FUC,
	}
}

func (router *FoodHandlers) Standby(ctx context.Context) {
	router.Router.GET("/food/view", func(ctx *gin.Context) {
		data, err := router.FUC.ViewAll(ctx)
		if err != nil {
			log.Fatal(err)
		}
		ctx.IndentedJSON(http.StatusOK, data)
	})

	router.Router.GET("/food/view/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}
		data, err := router.FUC.ViewOne(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		ctx.IndentedJSON(http.StatusOK, data)
	})

	router.Router.PATCH("/food/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		var status entity.Food
		ctx.BindJSON(&status)
		err := router.FUC.UpdateStatus(ctx, id, status)
		if err != nil {
			ctx.JSON(http.StatusForbidden, map[string]any{
				"message": "surprise its an error",
			})
			return
		}
		data, err := router.FUC.ViewOne(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		ctx.IndentedJSON(http.StatusOK, data)
	})
}
