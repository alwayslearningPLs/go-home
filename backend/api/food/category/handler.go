package category

import (
	"net/http"
	"strconv"

	"github.com/MrTimeout/go-home/backend/api/utils"
	"github.com/gin-gonic/gin"
)

const (
	// CategoriesPath retrieves all the categories inside the db
	CategoriesPath = "/food/categories"
	// CategoryByNamePath is used by the handler GetCategoryByName to
	// return the category by name.
	CategoryByNamePath = CategoriesPath + "/:" + CategoryNameParam

	// CategoryNameParam is the category name value
	CategoryNameParam = "category-name"
)

func GetCategories(c *gin.Context) {
	categories, err := getCategories(c.Request.Context(), utils.ParseRequest(c, FoodCategory{}))
	if err != nil {
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    categories,
	})
}

func GetCategoryByName(c *gin.Context) {
	category, err := getCategories(c.Request.Context(), utils.ParseRequest(c, FoodCategory{Name: c.Param(CategoryNameParam)}))
	if err != nil {
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    category,
	})
}

func AddCategory(c *gin.Context) {
	var category FoodCategory
	if err := c.Bind(&category); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	if _, err := addCategory(c.Request.Context(), &category); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    category,
	})
}

func DelCategory(c *gin.Context) {
	var (
		rows int64
		err  error
	)
	if rows, err = delCategory(c.Request.Context(), FoodCategory{Name: c.Param(CategoryNameParam)}); err != nil {
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
