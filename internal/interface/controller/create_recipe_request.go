package controller

type createRecipeRequest struct {
	Name string `json:"name" binding:"required"`
}
