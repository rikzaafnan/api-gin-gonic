package repository

import (
	"api-gin-gonic/models"
	"errors"

	"gorm.io/gorm"
)

type MovieRepositoryImplementation interface {
	Save(input models.Movie, tx *gorm.DB) (models.Movie, error)
	Update(movieId int, newInput models.Movie, tx *gorm.DB) (bool, error)
	FindById(movieId int, tx *gorm.DB) (models.Movie, error)
	FindAll(tx *gorm.DB) ([]models.Movie, error)
	Delete(movieId int, tx *gorm.DB) (bool, error)
}

type movieRepository struct {
}

func NewMovieRepository() *movieRepository {
	return &movieRepository{}
}

func (r *movieRepository) Save(input models.Movie, tx *gorm.DB) (models.Movie, error) {

	err := tx.Model(&input).Create(&input).Error
	if err != nil {
		// panic(err)
		return input, err
	}

	return input, nil

}

func (r *movieRepository) Update(movieId int, newInput models.Movie, tx *gorm.DB) (bool, error) {
	err := tx.Model(&newInput).Where("id = ?", movieId).Updates(&newInput).Error
	if err != nil {
		// panic(err)
		return false, err
	}

	return true, nil
}

func (r *movieRepository) FindById(movieId int, tx *gorm.DB) (models.Movie, error) {
	var movie models.Movie

	err := tx.Model(&movie).Where("id = ?", movieId).First(&movie).Error
	if err != nil {
		// panic(err)
		return movie, err
	}

	if movie.ID == 0 {
		return movie, errors.New("record not found")
	}

	return movie, nil
}

func (r *movieRepository) FindAll(tx *gorm.DB) ([]models.Movie, error) {
	var movies []models.Movie

	err := tx.Model(&movies).Find(&movies).Error
	if err != nil {
		// panic(err)
		return movies, err
	}

	return movies, nil
}

func (r *movieRepository) Delete(movieId int, tx *gorm.DB) (bool, error) {

	err := tx.Model(&models.Movie{}).Where("id = ?", movieId).Delete(&models.Movie{}).Error
	if err != nil {
		// panic(err)
		return false, err
	}

	return true, nil
}
