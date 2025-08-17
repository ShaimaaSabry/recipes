package controller

import (
	"bytes"
	"github.com/ShaimaaSabry/recipes/internal/domain/model"
	"github.com/ShaimaaSabry/recipes/internal/usecase/createrecipe"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRecipe_201(t *testing.T) {
	router, mockCreateRecipeInteractor := setupCreateRecipeInteractorMock()

	// given
	recipe := model.Of(1, "Mock Recipe")
	mockCreateRecipeInteractor.
		On("Execute", createrecipe.Command{Name: "Mock Recipe"}).
		Return(recipe, nil).
		Once()

	requestBody := bytes.NewBufferString(`{"name":"Mock Recipe"}`)
	request := httptest.NewRequest("POST", "/recipes", requestBody)

	// when
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// then
	require.Equal(t, http.StatusCreated, recorder.Code)
	require.Contains(t, recorder.Body.String(), `"Mock Recipe"`)
	mockCreateRecipeInteractor.AssertCalled(t, "Execute", createrecipe.Command{Name: "Mock Recipe"})
	mockCreateRecipeInteractor.AssertExpectations(t)
}

func TestCreateRecipe_400_EmptyName(t *testing.T) {
	router, mockCreateRecipeInteractor := setupCreateRecipeInteractorMock()

	// given
	requestBody := bytes.NewBufferString(`{"name":""}`)
	request := httptest.NewRequest("POST", "/recipes", requestBody)

	// when
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// then
	require.Equal(t, http.StatusBadRequest, recorder.Code)
	require.Contains(t, recorder.Body.String(), `"Name"`)
	mockCreateRecipeInteractor.AssertNotCalled(t, "Execute", mock.Anything)
}
