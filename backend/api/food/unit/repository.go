package unit

import (
	"context"
	"errors"

	sca "github.com/MrTimeout/go-home/backend/api/food/subcategory"
	"github.com/MrTimeout/go-home/backend/api/utils"
	"github.com/MrTimeout/go-home/backend/internals/config"
	"gorm.io/gorm"
)

func addUnit(ctx context.Context, fu *FoodUnit) error {
	return config.GetInstance(ctx).Transaction(func(tx *gorm.DB) error {
		txx := sca.WhereSubcategories(tx, fu.FoodSubcategory).Find(&fu.FoodSubcategory)
		if txx.Error != nil {
			return txx.Error
		} else if fu.FoodSubcategory.ID == 0 {
			return errors.New("joder bro")
		}

		return tx.Create(fu).Error
	})
}

func delUnit(ctx context.Context, fu FoodUnit) (rows int64, err error) {
	tx := WhereUnit(SubQueryUnit(config.GetInstance(ctx), fu), fu).Delete(&fu)
	return tx.RowsAffected, tx.Error
}

func getUnits(ctx context.Context, wrap utils.WrapperRequest[FoodUnit]) ([]FoodUnit, error) {
	var result []FoodUnit

	tx := WhereUnit(wrap.ToScope(config.GetInstance(ctx)), wrap.Body)

	if wrap.Body.FoodSubcategory.FoodCategory.Name != "" {
		tx = JoinCategories(tx, wrap.Body).Find(&result)
	} else {
		tx = JoinSubcategories(tx, wrap.Body).Find(&result)
	}

	return result, tx.Error
}

func WhereUnit(db *gorm.DB, fu FoodUnit) *gorm.DB {
	if fu.ID != 0 {
		db = db.Where(fu.TableName()+".id = ?", fu.ID)
	}

	if fu.Name != "" {
		db = db.Where(fu.TableName()+".name = ?", fu.Name)
	}

	if fu.Description != "" {
		db = db.Where(fu.TableName()+".description like ?", fu.Description)
	}

	return db
}

func JoinSubcategories(db *gorm.DB, fu FoodUnit) *gorm.DB {
	db = db.Joins("JOIN " + fu.FoodSubcategory.TableName() + " USING(food_subcategory_id)")
	return sca.WhereSubcategories(db, fu.FoodSubcategory)
}

func JoinCategories(db *gorm.DB, fu FoodUnit) *gorm.DB {
	return sca.JoinCategories(JoinSubcategories(db, fu), fu.FoodSubcategory)
}

func SubQueryUnit(db *gorm.DB, fu FoodUnit) *gorm.DB {
	return db.Where(fu.TableName()+".food_subcategory_id IN (?)", sca.SelectWhereSubcategories(db, fu.FoodSubcategory, "id"))
}
