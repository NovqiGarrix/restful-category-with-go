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
	"testing"
)

func TestWithNoDataFindAllCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories", nil)
	recorder := httptest.NewRecorder()

	router := test.SetupTest(db)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytesBody, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(bytesBody, &responseBody)

	assert.Equal(t, int(responseBody["code"].(float64)), http.StatusOK)
	assert.Equal(t, responseBody["status"], "OK")

	assert.Equal(t, len(responseBody["data"].([]interface{})), 0)

}

func TestWithSomeDataFindAllCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")

	validate := validator.New()

	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo, db, validate)

	var rawCategories = []web.CategoryCreateRequest{
		{
			Name: "Phone",
		},
		{
			Name: "Laptop",
		},
		{
			Name: "Tool",
		},
	}

	for _, rawCategory := range rawCategories {
		categoryService.Create(context.Background(), rawCategory)
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories", nil)
	recorder := httptest.NewRecorder()

	router := test.SetupTest(db)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytesBody, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(bytesBody, &responseBody)

	assert.Equal(t, int(responseBody["code"].(float64)), http.StatusOK)
	assert.Equal(t, responseBody["status"], "OK")

	resData := responseBody["data"].([]interface{})

	for i, category := range resData {
		c := category.(map[string]interface{})
		actual := rawCategories[i]

		assert.Equal(t, c["name"], actual.Name)
	}

	assert.Equal(t, len(resData), len(rawCategories))

}
