package database

import (
	"log"

	"golearn/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Lesson{},
		&models.Quiz{},
		&models.Question{},
		&models.Choice{},
		&models.QuizResult{},
		&models.QuizAnswer{},
		&models.Enrollment{},
		&models.Progress{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
	return db
}
