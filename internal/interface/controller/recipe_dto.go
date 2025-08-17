package controller

import (
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
)

type recipeDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func covertRecipeToDTO(recipe *model.Recipe) recipeDTO {
	return recipeDTO{
		Id:   recipe.Id(),
		Name: recipe.Name(),
	}
}
