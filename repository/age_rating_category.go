package repository

import (
	"api-gin-gonic/models"
	"errors"

	"gorm.io/gorm"
)

type AgeRatingCategoryRepository interface {
	Save(input models.AgeRatingCategory, tx *gorm.DB) (models.AgeRatingCategory, error)
	Update(ageRatingCategoryId int, newInput models.AgeRatingCategory, tx *gorm.DB) (bool, error)
	FindById(ageRatingCategoryId int, tx *gorm.DB) (models.AgeRatingCategory, error)
	FindAll(tx *gorm.DB) ([]models.AgeRatingCategory, error)
	Delete(ageRatingCategoryId int, tx *gorm.DB) (bool, error)
}

type repository struct {
}

func NewAgeRatingCategoryRepository() *repository {
	return &repository{}
}

func (r *repository) Save(input models.AgeRatingCategory, tx *gorm.DB) (models.AgeRatingCategory, error) {

	err := tx.Model(&input).Create(&input).Error
	if err != nil {
		// panic(err)
		return input, err
	}

	return input, nil

}

func (r *repository) Update(ageRatingCategoryId int, newInput models.AgeRatingCategory, tx *gorm.DB) (bool, error) {
	err := tx.Model(&newInput).Where("id = ?", ageRatingCategoryId).Updates(&newInput).Error
	if err != nil {
		// panic(err)
		return false, err
	}

	return true, nil
}

func (r *repository) FindById(ageRatingCategoryId int, tx *gorm.DB) (models.AgeRatingCategory, error) {
	var ageRatingCategory models.AgeRatingCategory

	err := tx.Model(&ageRatingCategory).Where("id = ?", ageRatingCategoryId).First(&ageRatingCategory).Error
	if err != nil {
		// panic(err)
		return ageRatingCategory, err
	}

	if ageRatingCategory.ID == 0 {
		return ageRatingCategory, errors.New("record not found")
	}

	return ageRatingCategory, nil
}

func (r *repository) FindAll(tx *gorm.DB) ([]models.AgeRatingCategory, error) {
	var ageRatingCategories []models.AgeRatingCategory

	err := tx.Model(&ageRatingCategories).Find(&ageRatingCategories).Error
	if err != nil {
		// panic(err)
		return ageRatingCategories, err
	}

	return ageRatingCategories, nil
}

func (r *repository) Delete(ageRatingCategoryId int, tx *gorm.DB) (bool, error) {

	err := tx.Model(&models.AgeRatingCategory{}).Where("id = ?", ageRatingCategoryId).Delete(&models.AgeRatingCategory{}).Error
	if err != nil {
		// panic(err)
		return false, err
	}

	return true, nil
}
