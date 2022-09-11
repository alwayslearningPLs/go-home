package unit

import (
	"errors"
	"net/http"
	"strconv"

	ca "github.com/MrTimeout/go-home/backend/api/food/category"
	sca "github.com/MrTimeout/go-home/backend/api/food/subcategory"
	"github.com/MrTimeout/go-home/backend/api/utils"
	"github.com/gin-gonic/gin"
)

const (
	// UnitPathName is the unit path name
	UnitPathName = "/units"

	// UnitsBySubcategoriesPath retrieves all the units inside a subcategory.
	// /food/subcategories/:subcategory-name/units
	UnitsBySubcategoriesPath = sca.SubcategoriesPathName + "/:" + sca.SubcategoryNameParam + UnitPathName
	// UnitBySubcategoryPath returns the unit by subcategory.
	// /food/subcategories/:subcategory-name/units/:unit-name
	UnitBySubcategoryPath = UnitsBySubcategoriesPath + "/:" + UnitNameParam

	// UnitsByCategoriesPath retrieves all the units inside a category.
	// /food/categories/:category-name/units
	UnitsByCategoriesPath = ca.CategoryByNamePath + UnitPathName
	// UnitByCategoriesPath retrieves the unit inside a category.
	// /food/categories/:category-name/units/:unit-name
	UnitByCategoriesPath = ca.CategoryByNamePath + UnitPathName + "/:" + UnitNameParam

	// UnitNameParam is the unit name param
	UnitNameParam = "unit-name"
)

// ErrUnitsNotFound is returned when not units were found.
var ErrUnitsNotFound = errors.New("units not found")

func GetUnitsBySubcategory(c *gin.Context) {
	units, err := getUnits(c.Request.Context(), utils.ParseRequest(c, newFoodUnitFromParams(c)))
	if err != nil || len(units) == 0 {
		if err == nil {
			err = ErrUnitsNotFound
		}
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    units,
	})
}

func GetUnitBySubcategory(c *gin.Context) {
	subcategory, err := getUnits(c.Request.Context(), utils.ParseRequest(c, newFoodUnitFromParams(c)))
	if err != nil || len(subcategory) == 0 {
		if err == nil {
			err = ErrUnitsNotFound
		}
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    subcategory,
	})
}

func GetUnitsByCategory(c *gin.Context) {
	units, err := getUnits(c.Request.Context(), utils.ParseRequest(c, newFoodUnitFromParams(c)))
	if err != nil || len(units) == 0 {
		if err == nil {
			err = ErrUnitsNotFound
		}
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    units,
	})
}

func GetUnitByCategory(c *gin.Context) {
	subcategory, err := getUnits(c.Request.Context(), utils.ParseRequest(c, newFoodUnitFromParams(c)))
	if err != nil || len(subcategory) == 0 {
		if err == nil {
			err = ErrUnitsNotFound
		}
		utils.ErrRes(c, err, http.StatusNotFound)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    subcategory,
	})
}

func AddUnit(c *gin.Context) {
	var unit FoodUnit
	if err := c.Bind(&unit); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	unit.FoodSubcategory.Name = c.Param(sca.SubcategoryNameParam)

	if err := addUnit(c.Request.Context(), &unit); err != nil {
		utils.ErrRes(c, err, http.StatusBadRequest)
		return
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: utils.Negotiate,
		Data:    unit,
	})
}

func DelUnit(c *gin.Context) {
	var (
		rows int64
		err  error
	)
	if rows, err = delUnit(c.Request.Context(), newFoodUnitFromParams(c)); err != nil {
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

func newFoodUnitFromParams(pParser utils.ParamParser) FoodUnit {
	return FoodUnit{
		Name: pParser.Param(UnitNameParam),
		FoodSubcategory: sca.FoodSubcategory{
			Name: pParser.Param(sca.SubcategoryNameParam),
			FoodCategory: ca.FoodCategory{
				Name: pParser.Param(ca.CategoryNameParam),
			},
		},
	}
}
