package model

import (
	"fmt"
	apperrors "github.com/ShaimaaSabry/recipes/internal/domain/errors"
)

var (
	ErrInvalidRecipeName = fmt.Errorf("%w: recipe name cannot be empty", apperrors.ErrInvalidInput)
)

type Recipe struct {
	id   int
	name string
}

func NewRecipe(name string) (*Recipe, error) {
	if name == "" {
		return nil, ErrInvalidRecipeName
	}

	return &Recipe{
		id:   0, // ID will be set by the repository
		name: name,
	}, nil
}

func Of(id int, name string) *Recipe {

	return &Recipe{
		id:   id,
		name: name,
	}
}

func (r Recipe) Id() int {
	return r.id
}

func (r Recipe) Name() string {
	return r.name
}
