package main

import (
	"api-gin-gonic/config"
	"api-gin-gonic/docs"
	"api-gin-gonic/repository"
	"api-gin-gonic/routes"
	"api-gin-gonic/service"
	"fmt"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {

	fmt.Println("coba online")

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectionDB()
	postgree, _ := db.DB()
	defer postgree.Close()

	ageRatingCategoryRepository := repository.NewAgeRatingCategoryRepository()
	ageRatingCategoryService := service.NewAgeRatingCategoryService(db, ageRatingCategoryRepository)

	r := routes.SetupRouter(ageRatingCategoryService)
	r.Run()

}
