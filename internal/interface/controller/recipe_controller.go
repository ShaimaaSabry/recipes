package controller

import (
	"fmt"
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
	"github.com/ShaimaaSabry/recipes/internal/usecase/createrecipe"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getRecipesInteractor interface {
	Execute() []model.Recipe
}

type createRecipeInteractor interface {
	Execute(command createrecipe.Command) (*model.Recipe, error)
}

type RecipeController struct {
	getRecipesInteractor   getRecipesInteractor
	createRecipeInteractor createRecipeInteractor
}

func NewRecipeController(
	getRecipesInteractor getRecipesInteractor,
	createRecipeInteractor createRecipeInteractor,
) *RecipeController {
	return &RecipeController{
		getRecipesInteractor:   getRecipesInteractor,
		createRecipeInteractor: createRecipeInteractor,
	}
}

// GetRecipes godoc
// @Summary     List recipes
// @Tags        recipes
// @Success     200 {array}  recipeDTO
// @Router      /recipes [get]
func (h *RecipeController) GetRecipes(ctx *gin.Context) {
	recipes := h.getRecipesInteractor.Execute()

	var dtos []recipeDTO
	for _, r := range recipes {
		dto := covertRecipeToDTO(&r)
		dtos = append(dtos, dto)
	}

	ctx.JSON(
		http.StatusOK,
		dtos,
	)
}

// CreateRecipe godoc
// @Summary     Create a new recipe
// @Tags        recipes
// @Param       payload body createRecipeRequest true "Create recipe payload"
// @Success     201 {object} recipeDTO
// @Router      /recipes [post]
func (h *RecipeController) CreateRecipe(ctx *gin.Context) {
	var request createRecipeRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(err)
		ctx.Error(err)
		return
	}

	recipe, err := h.createRecipeInteractor.Execute(createrecipe.Command{Name: request.Name})
	if err != nil {
		ctx.Error(err)
		return
	}

	dto := covertRecipeToDTO(recipe)

	ctx.JSON(
		http.StatusCreated,
		dto,
	)
}
