package main

import (
	"context"
	"time"

	ca "github.com/MrTimeout/go-home/backend/api/food/category"
	sca "github.com/MrTimeout/go-home/backend/api/food/subcategory"
	u "github.com/MrTimeout/go-home/backend/api/food/unit"
	"github.com/MrTimeout/go-home/backend/internals/cmd"
	"github.com/MrTimeout/go-home/backend/internals/config"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}

	ctx, cl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cl()

	config.GetInstance(ctx).AutoMigrate(&ca.FoodCategory{}, &sca.FoodSubcategory{}, &u.FoodUnit{})

	router := gin.New()

	food := router.Group("/food")
	{
		food.GET(ca.CategoriesPath, ca.GetCategories)
		food.POST(ca.CategoriesPath, ca.AddCategory)
		food.GET(ca.CategoryByNamePath, ca.GetCategoryByName)
		food.DELETE(ca.CategoryByNamePath, ca.DelCategory)

		food.GET(sca.SubcategoriesPath, sca.GetSubcategories)
		food.POST(sca.SubcategoriesPath, sca.AddSubcategory)
		food.GET(sca.SubcategoryByNamePath, sca.GetSubcategoryByName)
		food.DELETE(sca.SubcategoryByNamePath, sca.DelSubcategory)

		food.GET(u.UnitsBySubcategoriesPath, u.GetUnitsBySubcategory)
		food.POST(u.UnitsBySubcategoriesPath, u.AddUnit)
		food.GET(u.UnitBySubcategoryPath, u.GetUnitBySubcategory)
		food.DELETE(u.UnitBySubcategoryPath, u.DelUnit)

		food.GET(u.UnitsByCategoriesPath, u.GetUnitsByCategory)
		food.GET(u.UnitByCategoriesPath, u.GetUnitByCategory)
	}

	router.Run(":8080")
}
