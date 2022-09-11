package subcategory

import (
	"context"
	"errors"

	ca "github.com/MrTimeout/go-home/backend/api/food/category"
	"github.com/MrTimeout/go-home/backend/api/utils"
	"github.com/MrTimeout/go-home/backend/internals/config"
	"gorm.io/gorm"
)

func addSubcategory(ctx context.Context, fc *FoodSubcategory) error {
	return config.GetInstance(ctx).Transaction(func(tx *gorm.DB) error {
		txx := tx.Where(&fc.FoodCategory).Find(&fc.FoodCategory)
		if txx.Error != nil {
			return txx.Error
		} else if fc.FoodCategory.ID == 0 {
			return errors.New("joder bro")
		}

		return tx.Create(fc).Error
	})
}

func delSubcategory(ctx context.Context, fc FoodSubcategory) (rows int64, err error) {
	tx := WhereSubcategories(SubqueryCategories(config.GetInstance(ctx), fc), fc).Delete(fc)
	return tx.RowsAffected, tx.Error
}

func getSubcategories(ctx context.Context, wrap utils.WrapperRequest[FoodSubcategory]) ([]FoodSubcategory, error) {
	var result []FoodSubcategory

	tx := ca.WhereCategories(
		WhereSubcategories(wrap.ToScope(config.GetInstance(ctx)), wrap.Body).
			Joins("INNER JOIN food_categories USING(food_category_id)"), wrap.Body.FoodCategory).
		Find(&result)

	return result, tx.Error
}

func WhereSubcategories(db *gorm.DB, fc FoodSubcategory) *gorm.DB {
	if fc.ID != 0 {
		db = db.Where(fc.TableName()+".id = ?", fc.ID)
	}

	if fc.Name != "" {
		db = db.Where(fc.TableName()+".name = ?", fc.Name)
	}

	if fc.Description != "" {
		db = db.Where(fc.TableName()+".description like ?", fc.Description)
	}

	return db
}

func JoinCategories(db *gorm.DB, fs FoodSubcategory) *gorm.DB {
	return ca.WhereCategories(db.Joins("JOIN "+fs.FoodCategory.TableName()+" USING(food_category_id)"), fs.FoodCategory)
}

func SubqueryCategories(db *gorm.DB, fs FoodSubcategory) *gorm.DB {
	return db.Where(fs.TableName()+".food_category_id IN (?)", ca.SelectWhereCategories(db, fs.FoodCategory, "id"))
}

func SelectWhereSubcategories(db *gorm.DB, fs FoodSubcategory, projection ...string) *gorm.DB {
	return WhereSubcategories(db.Table(fs.TableName()).Select(projection), fs)
}
