package repository

import (
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
)

var db = []model.Recipe{}

type RecipeInMemoryRepository struct {
}

func (r *RecipeInMemoryRepository) SaveRecipe(recipe *model.Recipe) (*model.Recipe, error) {
	db = append(db, *recipe)
	return recipe, nil
}

func (r *RecipeInMemoryRepository) GetRecipes() []model.Recipe {
	return db
}
