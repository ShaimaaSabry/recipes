package controller

import (
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
	"github.com/ShaimaaSabry/recipes/internal/interface/controller/middleware"
	"github.com/ShaimaaSabry/recipes/internal/usecase/createrecipe"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type createRecipeInteractorMock struct{ mock.Mock }

func (m *createRecipeInteractorMock) Execute(cmd createrecipe.Command) (*model.Recipe, error) {
	args := m.Called(cmd)
	return args.Get(0).(*model.Recipe), args.Error(1)
}

type getRecipesInteractorMock struct{}

func (n getRecipesInteractorMock) Execute() []model.Recipe { return nil }

func setupRouter(controller RecipeController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	r.POST("/recipes", controller.CreateRecipe)
	r.GET("/recipes", controller.GetRecipes)
	return r
}

func setupCreateRecipeInteractorMock() (*gin.Engine, *createRecipeInteractorMock) {
	mockCreateRecipeInteractor := new(createRecipeInteractorMock)

	controller := NewRecipeController(&getRecipesInteractorMock{}, mockCreateRecipeInteractor)

	router := setupRouter(*controller)

	return router, mockCreateRecipeInteractor
}
