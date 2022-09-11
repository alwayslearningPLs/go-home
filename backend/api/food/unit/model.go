package unit

import (
	"encoding/xml"

	sca "github.com/MrTimeout/go-home/backend/api/food/subcategory"
)

// FoodUnit
//
// It is the representation of a piece of food. It doesn't contain any variety.
//
// swagger:model food-unit
type FoodUnit struct {
	// swagger:ignore
	XMLName xml.Name `gorm:"-" json:"-" xml:"FoodUnit"`
	// swagger:ignore
	ID int `gorm:"column:food_unit_id;primaryKey" json:"-" xml:"-"`
	// Name that uniquely identifies a piece of food
	//
	// required: true
	// min length: 2
	// example: banana
	Name string `gorm:"column:name;not null;unique" json:"name" xml:"Name"`
	// Description of the food unit. It can be as large as you want
	//
	// required: true
	// min length: 5
	// example: a long curved fruit which grows in clusters and has soft pulpy flesh and yellow skin when ripe.
	Description       string              `gorm:"column:description;not null" json:"description" xml:"Description"`
	FoodSubcategoryID int                 `json:"-" xml:"-"`
	FoodSubcategory   sca.FoodSubcategory `json:"-" xml:"-"`
}

// OrderByColumnsAllowed return the list of columns allowed to order by.
func (FoodUnit) OrderByColumnsAllowed() map[string]any {
	return map[string]any{"id": struct{}{}, "name": struct{}{}}
}

// TableName returns the name of the table that is going to be used to represent the FoodSubcategory struct
func (FoodUnit) TableName() string {
	return "food_units"
}
