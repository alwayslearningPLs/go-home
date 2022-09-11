package category

import (
	"context"

	"github.com/MrTimeout/go-home/backend/api/utils"
	"github.com/MrTimeout/go-home/backend/internals/config"
	"gorm.io/gorm"
)

func addCategory(ctx context.Context, fc *FoodCategory) (rows int64, err error) {
	tx := config.GetInstance(ctx).Create(fc)
	return tx.RowsAffected, tx.Error
}

func delCategory(ctx context.Context, fc FoodCategory) (rows int64, err error) {
	tx := config.GetInstance(ctx).Where(&fc).Delete(fc)
	return tx.RowsAffected, tx.Error
}

func getCategories(ctx context.Context, wrap utils.WrapperRequest[FoodCategory]) ([]FoodCategory, error) {
	var result []FoodCategory

	tx := wrap.ToScope(config.GetInstance(ctx).Where(wrap.Body)).Find(&result)

	return result, tx.Error
}

func WhereCategories(db *gorm.DB, fc FoodCategory) *gorm.DB {
	if fc.ID != 0 {
		db = db.Where(fc.TableName()+".id = ?", fc.ID)
	}

	if fc.Description != "" {
		db = db.Where(fc.TableName()+".description like ?", fc.Description)
	}

	if fc.Name != "" {
		db = db.Where(fc.TableName()+".name = ?", fc.Name)
	}

	return db
}

func SelectWhereCategories(db *gorm.DB, fc FoodCategory, projection ...string) *gorm.DB {
	return db.Table(fc.TableName()).Select(projection).Where(&fc)
}
