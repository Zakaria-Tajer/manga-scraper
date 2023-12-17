package db

import (
	"client/consumer/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Connection() *gorm.DB {

	dsn := "host=localhost user=root password=root dbname=consumer port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection failed")
	} else {
		log.Println("connection success")
	}

	dbInstance = db
	err = db.AutoMigrate(&models.Manga{}, &models.MangaDetails{}, &models.MangaChapters{})

	if err != nil {
		log.Fatal("error migrating database: ", err)
	}

	log.Println("Migration succeded")
	return db
}

func GetDB() *gorm.DB {

	return dbInstance
}
