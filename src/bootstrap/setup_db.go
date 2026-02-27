package bootstrap

import (
	"desent/src/migrations"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	log.Printf("Connecting to in-memory database SQLite")
	db, err := gorm.Open(
		sqlite.Open("sqlite.db"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(migrations.MigrationLists...); err != nil {
		panic(err)
	}

	return db
}
