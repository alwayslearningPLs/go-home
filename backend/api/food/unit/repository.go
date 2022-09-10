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

func delSubCategory(ctx context.Context, fu FoodUnit) (rows int64, err error) {
	tx := JoinUnit(sca.WhereSubcategories(config.GetInstance(ctx), fu.FoodSubcategory), fu).Delete(&fu)
	return tx.RowsAffected, tx.Error
}

func getUnits(ctx context.Context, wrap utils.WrapperRequest[FoodUnit]) ([]FoodUnit, error) {
	var result []FoodUnit

	tx := WhereUnit(wrap.ToScope(config.GetInstance(ctx)), wrap.Body).Preload("FoodSubcategory").Find(&result)

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

func JoinUnit(db *gorm.DB, fu FoodUnit) *gorm.DB {
	db = db.Joins("JOIN " + fu.FoodSubcategory.TableName() + " ON " + fu.FoodSubcategory.TableName() + ".id = " + fu.TableName() + ".id")
	return WhereUnit(db, fu)
}

func doGetUnit(ctx context.Context, wrap utils.WrapperRequest[FoodUnit]) *gorm.DB {
	return wrap.ToScope(config.GetInstance(ctx)).
		Where("food_subcategories.name = ?", wrap.Body.FoodSubcategory.Name)
}
