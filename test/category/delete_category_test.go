package category

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"restful-api/model/domain/web"
	"restful-api/repository"
	"restful-api/service"
	"restful-api/test"
	"strconv"
	"testing"
)

func TestSuccessDeleteCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")

	validate := validator.New()

	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo, db, validate)

	newCategory := web.CategoryCreateRequest{Name: "Macbook"}

	newCategoryResponse := categoryService.Create(context.Background(), newCategory)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:4000/api/categories/"+strconv.Itoa(newCategoryResponse.Id), nil)
	recorder := httptest.NewRecorder()

	router := test.SetupTest(db)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bytesBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusOK)

	var responseBody map[string]interface{}
	_ = json.Unmarshal(bytesBody, &responseBody)

	assert.Equal(t, int(responseBody["code"].(float64)), http.StatusOK)
	assert.Equal(t, responseBody["status"], "OK")

	assert.Equal(t, int(responseBody["data"].(float64)), newCategoryResponse.Id)

	// Check if the data is deleted
	panicFunc := func() {
		categoryService.FindById(context.Background(), newCategoryResponse.Id)
	}

	assert.Panics(t, panicFunc)

}

func TestNotFoundDeleteCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")
	
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:4000/api/categories/000", nil)
	recorder := httptest.NewRecorder()

	router := test.SetupTest(db)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bytesBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusNotFound)

	var responseBody map[string]interface{}
	_ = json.Unmarshal(bytesBody, &responseBody)

	assert.Equal(t, int(responseBody["code"].(float64)), http.StatusNotFound)
	assert.Equal(t, responseBody["status"], "NOT FOUND")

	assert.NotEqual(t, responseBody["error"], nil)

}
