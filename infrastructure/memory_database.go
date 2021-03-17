package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MemoryDatabase(cbs ...func(db *gorm.DB)) *gorm.DB {
	dsn := "host=localhost user=postgres password=secret dbname=pismo port=5432 sslmode=disable search_path=pg_temp"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = DatabaseAutoMigrate(db); err != nil {
		panic(err)
	}

	if err = DatabaseSeed(db); err != nil {
		panic(err)
	}

	for _, cb := range cbs {
		cb(db)
	}

	return db
}
