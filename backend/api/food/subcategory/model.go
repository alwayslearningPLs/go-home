package subcategory

import (
	"encoding/xml"

	c "github.com/MrTimeout/go-home/backend/api/food/category"
)

// FoodSubcategory
//
// It contains all the subcategories for each category of food. It is another level
// of categorizing food.
//
// swagger:model food-subcategory
type FoodSubcategory struct {
	// swagger:ignore
	XMLName xml.Name `gorm:"-" json:"-" xml:"SubCategory"`
	// swagger:ignore
	ID int `gorm:"column:food_subcategory_id;primaryKey" json:"-" xml:"-"`
	// Name represents the name of the subcategory and is unique along the table
	//
	// required: true
	// min length: 2
	// example: Dark Green vegetable
	Name string `gorm:"column:name;not null;unique" json:"name,omitempty" xml:"Name"`
	// Description represents a little definition of each subcategory
	//
	// required: true
	// min length: 5
	// example: dark green vegetables as broccoli, collard greens, spinach, romaine, etc.
	Description    string         `gorm:"column:description;not null" json:"description,omitempty" xml:"Description"`
	FoodCategoryID int            `json:"-" xml:"-"`
	FoodCategory   c.FoodCategory `json:"-" xml:"-"`
}

// OrderByColumnsAllowed return the list of columns allowed to order by.
func (FoodSubcategory) OrderByColumnsAllowed() map[string]any {
	return map[string]any{"id": struct{}{}, "name": struct{}{}}
}

// TableName returns the name of the table that is going to be used to represent the FoodSubcategory struct
func (FoodSubcategory) TableName() string {
	return "food_subcategories"
}
