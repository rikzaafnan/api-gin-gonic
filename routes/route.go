package routes

import (
	"api-gin-gonic/controller"
	"api-gin-gonic/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(ageRatingCategoryService service.AgeRatingCategoryServiceImplementation) *gin.Engine {

	ageRatingCategoryController := controller.NewAgeRatingCategoryController(ageRatingCategoryService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "sukses")
	})

	r.GET("/age-rating-categories", ageRatingCategoryController.Index)
	r.POST("/age-rating-categories", ageRatingCategoryController.Store)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
