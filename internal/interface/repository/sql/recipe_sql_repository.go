package sql

import (
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
	"gorm.io/gorm"
)

type RecipeSqlRepository struct {
	db *gorm.DB
}

func NewRecipeSqlRepository(db *gorm.DB) *RecipeSqlRepository {
	err := db.AutoMigrate(&recipeDBO{})
	if err != nil {
		panic(err)
	}
	return &RecipeSqlRepository{db: db}
}

func (r *RecipeSqlRepository) SaveRecipe(recipe *model.Recipe) (*model.Recipe, error) {
	dbo := recipeDBO{
		Id:   recipe.Id(),
		Name: recipe.Name(),
	}

	tx := r.db.Save(&dbo)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return model.Of(dbo.Id, dbo.Name), nil
}

func (r *RecipeSqlRepository) GetRecipes() []model.Recipe {
	var out []recipeDBO

	tx := r.db.Find(&out)
	if tx.Error != nil {
		// Your signature doesn't return error, so return an empty slice
		panic(tx.Error)
	}

	var recipes []model.Recipe
	for _, dbo := range out {
		recipe := model.Of(dbo.Id, dbo.Name)
		recipes = append(recipes, *recipe)
	}
	return recipes
}
