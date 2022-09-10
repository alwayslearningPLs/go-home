package utils

import "gorm.io/gorm"

type RepositoryBuilder interface {
	Build() *gorm.DB
}

type OrderByAllower interface {
	OrderByColumnsAllowed() map[string]any
}

func (w WrapperRequest[T]) ToScope(db *gorm.DB) *gorm.DB {
	return w.limit(w.skip(w.orderBy(db)))
}

func (w WrapperRequest[T]) orderBy(db *gorm.DB) *gorm.DB {
	var cols = w.Body.OrderByColumnsAllowed()

	for _, each := range w.OrderBy {
		if _, ok := cols[each.Field]; ok {
			db = db.Order(each.String())
		}
	}

	return db
}

func (w WrapperRequest[T]) limit(db *gorm.DB) *gorm.DB {
	return db.Limit(w.Limit)
}

func (w WrapperRequest[T]) skip(db *gorm.DB) *gorm.DB {
	return db.Offset(w.Skip)
}
