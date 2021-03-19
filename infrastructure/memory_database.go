package infrastructure

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MemoryDatabase(cbs ...func(db *gorm.DB)) *gorm.DB {
	dsn := os.Getenv("TEST_DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = DatabaseAutoMigrate(db); err != nil {
		panic(err)
	}

	for _, cb := range cbs {
		cb(db)
	}

	return db
}
