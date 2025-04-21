package repository

import (
	"backend/internal/model/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB(connectionString string) {
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.Migrator().DropTable(&model.User{}) // Rmb to remove!
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func CloseDB() {
	db, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	db.Close()
}
