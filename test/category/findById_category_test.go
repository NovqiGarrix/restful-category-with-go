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

func TestSuccessFindByIdCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")

	validate := validator.New()

	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo, db, validate)

	newCategory := web.CategoryCreateRequest{
		Name: "Macbook",
	}

	// Create a new category first
	categoryResponse := categoryService.Create(context.Background(), newCategory)

	router := test.SetupTest(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories/"+strconv.Itoa(categoryResponse.Id), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, response.StatusCode, http.StatusOK)

	bytes, _ := io.ReadAll(response.Body)

	var resBody map[string]interface{}
	_ = json.Unmarshal(bytes, &resBody)

	assert.Equal(t, int(resBody["code"].(float64)), http.StatusOK)
	assert.Equal(t, resBody["status"], "OK")

	resData := resBody["data"].(map[string]interface{})
	assert.Equal(t, resData["name"], "Macbook")

}

func TestNotFoundFindByIdCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")
	router := test.SetupTest(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories/000", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, response.StatusCode, http.StatusNotFound)

	bytes, _ := io.ReadAll(response.Body)

	var resBody map[string]interface{}
	_ = json.Unmarshal(bytes, &resBody)

	assert.Equal(t, int(resBody["code"].(float64)), http.StatusNotFound)
	assert.Equal(t, resBody["status"], "NOT FOUND")

	assert.Equal(t, resBody["error"], "category not found")

}
