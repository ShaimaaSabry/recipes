package getrecipes

import (
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
)

type getRecipesInteractor interface {
	GetRecipes() []model.Recipe
}

type Interactor struct {
	getRecipesInteractor getRecipesInteractor
}

func NewInteractor(
	getRecipesInteractor getRecipesInteractor,
) *Interactor {
	return &Interactor{
		getRecipesInteractor: getRecipesInteractor,
	}
}

func (c *Interactor) Execute() []model.Recipe {
	return c.getRecipesInteractor.GetRecipes()
}
