package category

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"restful-api/test"
	"strings"
	"testing"
)

func TestSuccessCreateNewCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")
	router := test.SetupTest(db)

	requestBody := strings.NewReader(`
		{
			"name": "Phone"
		}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:4000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, response.StatusCode, http.StatusCreated)

	bytes, err := io.ReadAll(response.Body)
	assert.Equal(t, err, nil)

	var resBody map[string]interface{}
	err = json.Unmarshal(bytes, &resBody)
	assert.Equal(t, err, nil)

	assert.Equal(t, int(resBody["code"].(float64)), http.StatusCreated)
	assert.Equal(t, resBody["status"], "CREATED")

	resData := resBody["data"].(map[string]interface{})

	assert.Equal(t, resData["name"], "Phone")

}

func TestValidationFailedCreateNewCategory(t *testing.T) {

	db := test.SetupDB()
	test.TruncateDB(db, "category")
	router := test.SetupTest(db)

	requestBody := strings.NewReader(`
		{
			"name": ""
		}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:4000/api/categories", requestBody)
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
