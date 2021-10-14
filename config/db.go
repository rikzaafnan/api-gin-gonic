package config

import (
	"api-gin-gonic/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	dsn := "host=ec2-34-203-91-150.compute-1.amazonaws.com user=xrcumwrqnbzdhv password=789b348a6e4e5a586d84ef319014ba26e2beaa2882d6c41e4a1a8bc6d970d7a8 dbname=d6u51rvkrel1kt port=5432 "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.AgeRatingCategory{}, &models.Movie{})

	return db
}
