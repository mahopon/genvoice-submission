package repository

import (
	"backend/internal/model"
	"backend/internal/util"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(connectionString string) {
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop tables (for development purpose, remove them in production)
	db.Migrator().DropTable(&model.User{})     // Rmb to remove!
	db.Migrator().DropTable(&model.Survey{})   // Rmb to remove!
	db.Migrator().DropTable(&model.Answer{})   // Rmb to remove!
	db.Migrator().DropTable(&model.Question{}) // Rmb to remove!

	// Migrate the schema
	err = db.AutoMigrate(&model.User{}, &model.Survey{}, &model.Answer{}, &model.Question{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Insert default users if they don't exist
	pass1, hash1, _ := util.GenerateFromPassword("Userpassword123!")
	pass2, hash2, _ := util.GenerateFromPassword("Userpassword123!")
	pass3, hash3, _ := util.GenerateFromPassword("Userpassword123!")
	pass4, hash4, _ := util.GenerateFromPassword("adminpassword123!")
	users := []model.User{
		{Name: "User One", Username: "user1", Password: hash1 + ":" + pass1, Role: "USER"},
		{Name: "User Two", Username: "user2", Password: hash2 + ":" + pass2, Role: "USER"},
		{Name: "User Three", Username: "user3", Password: hash3 + ":" + pass3, Role: "USER"},
		{Name: "Admin User", Username: "admin", Password: hash4 + ":" + pass4, Role: "ADMIN"},
	}

	for _, user := range users {
		// Check if user already exists, if not insert them
		var count int64
		db.Model(&model.User{}).Where("username = ?", user.Username).Count(&count)
		if count == 0 {
			err := db.Create(&user).Error
			if err != nil {
				log.Printf("Failed to insert user %s: %v", user.Username, err)
			}
		}
	}
}

func CloseDB() {
	dbConn, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	dbConn.Close()
}
