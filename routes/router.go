package routes

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restful-api/exception"
)

func InitAllRoutes(db *sql.DB, validate *validator.Validate) *httprouter.Router {
	router := httprouter.New()

	// Category Routes
	InitCategoryRoutes(router, db, validate)

	router.GET("/health", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Println(request.Header.Get("Content-Type"))
		writer.WriteHeader(200)
		fmt.Fprint(writer, "I am Healthy, ready to go!!")
	})

	router.PanicHandler = exception.PanicHandler

	return router
}
