package config

import (
	"context"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// ConfigureDB is going to configure the database with the string passed as a parameter
func ConfigureDB(connStr string) {
	dbOnce.Do(func() {
		var err error

		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
}

func GetInstance(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func GetDrySession(ctx context.Context) *gorm.DB {
	return db.Session(&gorm.Session{DryRun: true}).WithContext(ctx)
}
