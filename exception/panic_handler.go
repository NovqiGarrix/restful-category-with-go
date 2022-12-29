package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"restful-api/helper"
	"restful-api/model/domain/web"
)

func PanicHandler(w http.ResponseWriter, request *http.Request, error interface{}) {

	if notFoundError(w, request, error) {
		return
	}
	if validationError(w, request, error) {
		return
	}
	
	internalServerError(w, request, error)

}

func validationError(w http.ResponseWriter, _ *http.Request, error interface{}) bool {

	exception, ok := error.(validator.ValidationErrors)

	if ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		response := web.ErrorResponseBody{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Error:  exception.Error(),
		}

		helper.EncodeResponseBody(w, response)
	}

	return ok

}

func notFoundError(w http.ResponseWriter, _ *http.Request, error interface{}) bool {

	exception, ok := error.(NotFoundException)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := web.ErrorResponseBody{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Error:  exception.Error,
		}

		helper.EncodeResponseBody(w, response)
	}

	return ok

}

func internalServerError(w http.ResponseWriter, _ *http.Request, err interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	var errorMsg string

	exception, ok := err.(error)
	if ok {
		errorMsg = exception.Error()
	} else {
		errorMsg = "Internal Server Error"
	}

	response := web.ErrorResponseBody{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  errorMsg,
	}

	helper.EncodeResponseBody(w, response)

}
