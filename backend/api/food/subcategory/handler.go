package subcategory

import (
	"errors"
	"net/http"
	"strconv"

	ca "github.com/MrTimeout/go-home/backend/api/food/category"
	"github.com/MrTimeout/go-home/backend/api/utils"
	"github.com/gin-gonic/gin"
)

const (
	// SubcategoriesPath retrieves all the subcategories inside the db by category-name
	SubcategoriesPath = ca.CategoryByNamePath + "/subcategories"
	// SubcategoryByNamePath is used by the handler GetCategoryByName to
	// return the category by name.
	SubcategoryByNamePath = SubcategoriesPath + "/:" + SubcategoryNameParam

	// SubcategoryNameParam is the category name value
	SubcategoryNameParam = "subcategory-name"
)

var ErrSubcategoryNotFound = errors.New("subcategory not found")

func GetSubcategories(c *gin.Context) {
	subcategories, err := getSubcategories(
		c.Request.Context(),
		utils.ParseRequest(c, FoodSubcategory{FoodCategory: ca.FoodCategory{Name: c.Param(ca.CategoryNameParam)}}),
	)
	if err != nil || len(subcategories) == 0 {
		if err == nil {
			err = ErrSubcategoryNotFound
		}
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    subcategories,
	})
}

func GetSubcategoryByName(c *gin.Context) {
	subcategory, err := getSubcategories(
		c.Request.Context(),
		utils.ParseRequest(c, newFoodSubcategoryFromParams(c)))
	if err != nil || len(subcategory) == 0 {
		if err == nil {
			err = ErrSubcategoryNotFound
		}
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    subcategory,
	})
}

func AddSubcategory(c *gin.Context) {
	var subcategory FoodSubcategory
	if err := c.Bind(&subcategory); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	subcategory.FoodCategory.Name = c.Param(ca.CategoryNameParam)

	if err := addSubcategory(c.Request.Context(), &subcategory); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    subcategory,
	})
}

func DelSubcategory(c *gin.Context) {
	var (
		rows int64
		err  error
	)
	if rows, err = delSubcategory(c.Request.Context(), newFoodSubcategoryFromParams(c)); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data: utils.WrapperResponse{
			Msg:  "category rows deleted " + strconv.Itoa(int(rows)),
			Code: http.StatusOK,
		},
	})
}

func newFoodSubcategoryFromParams(pParser utils.ParamParser) FoodSubcategory {
	return FoodSubcategory{
		Name:         pParser.Param(SubcategoryNameParam),
		FoodCategory: ca.FoodCategory{Name: pParser.Param(ca.CategoryNameParam)},
	}
}
