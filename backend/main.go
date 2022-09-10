package main

import (
	"context"
	"time"

	c "github.com/MrTimeout/go-home/backend/api/food/category"
	s "github.com/MrTimeout/go-home/backend/api/food/subcategory"
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

	config.GetInstance(ctx).AutoMigrate(&c.FoodCategory{}, &s.FoodSubcategory{})

	router := gin.New()

	router.GET(c.CategoriesPath, c.GetCategories)
	router.POST(c.CategoriesPath, c.AddCategory)
	router.GET(c.CategoryByNamePath, c.GetCategoryByName)
	router.DELETE(c.CategoryByNamePath, c.DelCategory)

	router.GET(s.SubcategoriesPath, s.GetSubcategories)
	router.POST(s.SubcategoriesPath, s.AddSubcategory)
	router.GET(s.SubcategoryByNamePath, s.GetSubcategoryByName)
	router.DELETE(s.SubcategoryByNamePath, s.DelSubcategory)

	router.Run(":8080")
}
