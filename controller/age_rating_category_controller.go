package controller

import (
	"api-gin-gonic/models"
	"api-gin-gonic/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ageRatingCategoryController struct {
	ageRatingCategoryService service.AgeRatingCategoryServiceImplementation
}

func NewAgeRatingCategoryController(ageRatingCategoryService service.AgeRatingCategoryServiceImplementation) *ageRatingCategoryController {

	return &ageRatingCategoryController{
		ageRatingCategoryService: ageRatingCategoryService,
	}

}

// GetAllAgeRatingCategory godoc
// @Summary Get all AgeRatingCategory.
// @Description Get a list of AgeRatingCategory.
// @Tags AgeRatingCategory
// @Produce json
// @Success 200 {object} []models.AgeRatingCategory
// @Router /age-rating-categories [get]
func (controller *ageRatingCategoryController) Index(c *gin.Context) {

	var ratings []models.AgeRatingCategory

	ratings, err := controller.ageRatingCategoryService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, gin.H{"data": ratings})

}

// CreateAgeRatingCategory godoc
// @Summary Create New AgeRatingCategory.
// @Description Creating a new AgeRatingCategory.
// @Tags AgeRatingCategory
// @Param Body body models.AgeRatingCategoryInput true "the body to create a new AgeRatingCategory"
// @Produce json
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories [post]
func (controller *ageRatingCategoryController) Store(c *gin.Context) {

	// Validate input
	var input models.AgeRatingCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ratings, err := controller.ageRatingCategoryService.Save(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, gin.H{"data": ratings})

}
