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
	"strings"
	"testing"
)

func TestSuccessUpdateNewCategory(t *testing.T) {

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

	requestBody := strings.NewReader(`
		{
			"name": "Gadget"
		}
	`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/categories/"+strconv.Itoa(categoryResponse.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, response.StatusCode, http.StatusOK)

	bytes, err := io.ReadAll(response.Body)
	assert.Equal(t, err, nil)

	var resBody map[string]interface{}
	err = json.Unmarshal(bytes, &resBody)
	assert.Equal(t, err, nil)

	assert.Equal(t, int(resBody["code"].(float64)), http.StatusOK)
	assert.Equal(t, resBody["status"], "OK")

	resData := resBody["data"].(map[string]interface{})

	assert.Equal(t, resData["name"], "Gadget")

}

func TestValidationFailedUpdateNewCategory(t *testing.T) {

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

	requestBody := strings.NewReader(`
		{
			"name": ""
		}
	`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/categories/"+strconv.Itoa(categoryResponse.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, response.StatusCode, http.StatusBadRequest)

	bytes, err := io.ReadAll(response.Body)
	assert.Equal(t, err, nil)

	var resBody map[string]interface{}
	err = json.Unmarshal(bytes, &resBody)
	assert.Equal(t, err, nil)

	assert.Equal(t, int(resBody["code"].(float64)), http.StatusBadRequest)
	assert.Equal(t, resBody["status"], "BAD REQUEST")

	assert.NotEqual(t, resBody["error"], nil)

}
