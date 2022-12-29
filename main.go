package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"restful-api/app"
	"restful-api/helper"
	"restful-api/routes"
)

func main() {

	db := app.GetDB()
	validate := validator.New()

	router := routes.InitAllRoutes(db, validate)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
