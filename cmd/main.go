package main

import (
	"github.com/ShaimaaSabry/recipes/internal/interface/controller"
	"github.com/ShaimaaSabry/recipes/internal/interface/controller/middleware"
	"github.com/ShaimaaSabry/recipes/internal/interface/repository/sql"
	"github.com/ShaimaaSabry/recipes/internal/usecase/createrecipe"
	"github.com/ShaimaaSabry/recipes/internal/usecase/getrecipes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// This import is REQUIRED so the generated docs package registers your spec.
	_ "github.com/ShaimaaSabry/recipes/docs"
)

// @title        Recipes API
// @version      1.0
// @description  Simple recipes service with CRUD endpoints.
func main() {
	recipeController := dependencyInjection()

	startServer(recipeController)
}

func dependencyInjection() *controller.RecipeController {
	//recipeRepository := &repository.RecipeInMemoryRepository{}

	dsn := "host=localhost user=postgres password=postgres dbname=recipes port=4444 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	recipeRepository := sql.NewRecipeSqlRepository(db)

	getRecipesInteractor := getrecipes.NewInteractor(recipeRepository)
	createRecipeInteractor := createrecipe.NewInteractor(recipeRepository)

	return controller.NewRecipeController(
		getRecipesInteractor,
		createRecipeInteractor,
	)
}

func startServer(recipeController *controller.RecipeController) {
	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())

	r.POST("/v1/recipes", recipeController.CreateRecipe)
	r.GET("/v1/recipes", recipeController.GetRecipes)

	setupSwagger(r)

	r.Run()
}

func setupSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // UI
}
