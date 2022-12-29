package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restful-api/helper"
	"restful-api/model/domain/web"
	"restful-api/service"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	requestBody := web.CategoryCreateRequest{}
	helper.DecodeRequestBody(request, &requestBody)

	categoryResponse := controller.CategoryService.Create(request.Context(), requestBody)

	writer.WriteHeader(http.StatusCreated)
	responseBody := web.ResponseBody{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   categoryResponse,
	}

	helper.EncodeResponseBody(writer, &responseBody)

}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	requestBody := web.CategoryUpdateRequest{}
	helper.DecodeRequestBody(request, &requestBody)

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	requestBody.Id = categoryId

	categoryResponse := controller.CategoryService.Update(request.Context(), requestBody)

	resBody := web.ResponseBody{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.EncodeResponseBody(writer, &resBody)

}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), categoryId)

	resBody := web.ResponseBody{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.EncodeResponseBody(writer, resBody)

}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	paramId := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(paramId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), categoryId)

	resBody := web.ResponseBody{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryId,
	}

	helper.EncodeResponseBody(writer, resBody)

}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	categories := controller.CategoryService.FindAll(request.Context())

	var data any
	if categories == nil {
		data = make([]string, 0)
	} else {
		data = categories
	}

	resBody := web.ResponseBody{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	helper.EncodeResponseBody(writer, resBody)

}
