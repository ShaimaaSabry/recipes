package createrecipe

import (
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
)

type recipeRepository interface {
	SaveRecipe(recipe *model.Recipe) (*model.Recipe, error)
}

type Interactor struct {
	recipeRepository recipeRepository
}

func NewInteractor(
	recipeRepository recipeRepository,
) *Interactor {
	return &Interactor{
		recipeRepository: recipeRepository,
	}
}

type Command struct {
	Name string
}

func (c *Interactor) Execute(command Command) (*model.Recipe, error) {
	recipe, err := model.NewRecipe(command.Name)
	if err != nil {
		return nil, err
	}

	recipe, err = c.recipeRepository.SaveRecipe(recipe)

	return recipe, err
}
