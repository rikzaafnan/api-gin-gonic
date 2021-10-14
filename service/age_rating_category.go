package service

import (
	"api-gin-gonic/models"
	"api-gin-gonic/repository"
	"time"

	"gorm.io/gorm"
)

type AgeRatingCategoryServiceImplementation interface {
	Save(input models.AgeRatingCategoryInput) (models.AgeRatingCategory, error)
	Update(ageRatingCategoryId int, newInput models.AgeRatingCategoryInput) (models.AgeRatingCategory, error)
	FindById(ageRatingCategoryId int) (models.AgeRatingCategory, error)
	FindAll() ([]models.AgeRatingCategory, error)
	Delete(ageRatingCategoryId int) (models.AgeRatingCategory, error)
}

type ageRatingCategoryService struct {
	db                          *gorm.DB
	ageRatingCategoryRepository repository.AgeRatingCategoryRepository
}

func NewAgeRatingCategoryService(db *gorm.DB, ageRatingCategoryRepository repository.AgeRatingCategoryRepository) *ageRatingCategoryService {
	return &ageRatingCategoryService{
		db:                          db,
		ageRatingCategoryRepository: ageRatingCategoryRepository,
	}
}

func (s *ageRatingCategoryService) Save(input models.AgeRatingCategoryInput) (models.AgeRatingCategory, error) {

	// var ageRatingCategory models.AgeRatingCategory

	tx := s.db.Begin()

	rating := models.AgeRatingCategory{Name: input.Name, Description: input.Description}

	ageRatingCategory, err := s.ageRatingCategoryRepository.Save(rating, tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}

	ageRatingCategory, err = s.ageRatingCategoryRepository.FindById(int(ageRatingCategory.ID), tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}
	tx.Commit()

	return ageRatingCategory, nil

}

func (s *ageRatingCategoryService) Update(ageRatingCategoryId int, newInput models.AgeRatingCategoryInput) (models.AgeRatingCategory, error) {

	tx := s.db.Begin()

	ageRatingCategory, err := s.ageRatingCategoryRepository.FindById(ageRatingCategoryId, tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}

	// rating := models.AgeRatingCategory{Name: newInput.Name, Description: newInput.Description}
	ageRatingCategory.Name = newInput.Name
	ageRatingCategory.Description = newInput.Description
	ageRatingCategory.UpdatedAt = time.Now()

	_, err = s.ageRatingCategoryRepository.Update(ageRatingCategoryId, ageRatingCategory, tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}

	ageRatingCategory, err = s.ageRatingCategoryRepository.FindById(int(ageRatingCategory.ID), tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}
	tx.Commit()

	return ageRatingCategory, nil

}

func (s *ageRatingCategoryService) FindById(ageRatingCategoryId int) (models.AgeRatingCategory, error) {

	tx := s.db.Begin()

	ageRatingCategory, err := s.ageRatingCategoryRepository.FindById(ageRatingCategoryId, tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}

	tx.Commit()

	return ageRatingCategory, nil

}

func (s *ageRatingCategoryService) FindAll() ([]models.AgeRatingCategory, error) {

	tx := s.db.Begin()

	// var ageRatingCategories []models.AgeRatingCategory

	ageRatingCategories, err := s.ageRatingCategoryRepository.FindAll(tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategories, err
	}

	tx.Commit()

	return ageRatingCategories, nil

}

func (s *ageRatingCategoryService) Delete(ageRatingCategoryId int) (models.AgeRatingCategory, error) {

	tx := s.db.Begin()

	ageRatingCategory, err := s.ageRatingCategoryRepository.FindById(ageRatingCategoryId, tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}

	_, err = s.ageRatingCategoryRepository.Delete(ageRatingCategoryId, tx)
	if err != nil {
		// panic(err)
		tx.Rollback()
		return ageRatingCategory, err
	}

	tx.Commit()

	return ageRatingCategory, nil

}
