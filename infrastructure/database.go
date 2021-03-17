package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=localhost user=postgres password=secret dbname=pismo port=5432 sslmode=disable"
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

	return db
}
