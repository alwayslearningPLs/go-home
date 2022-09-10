package category

import "encoding/xml"

// FoodCategory
//
// It contains all the categories to classify the food into big groups.
//
// swagger:model category
type FoodCategory struct {
	// swagger:ignore
	XMLName xml.Name `gorm:"-" json:"-" xml:"FoodCategory"`
	// swagger:ignore
	ID int `gorm:"column:id;primaryKey" json:"-" xml:"-"`
	// The name of the category
	//
	// required: true
	// min length: 2
	// example: grains
	Name string `gorm:"column:name;not null;unique" json:"name" xml:"Name"`
	// The description of the category. It should not be so long.
	//
	// required: true
	// min length: 10
	// example: a single fruit or seed of a cereal
	Description string `gorm:"column:description;not null" json:"description" xml:"Description"`
}

// OrderByColumnsAllowed will return the list of columns allowed to order by.
func (FoodCategory) OrderByColumnsAllowed() map[string]any {
	return map[string]any{"id": struct{}{}, "name": struct{}{}}
}

// TableName returns the name of table inside of the database.
func (FoodCategory) TableName() string {
	return "food_categories"
}
