package test

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"restful-api/app"
	"restful-api/helper"
	"restful-api/routes"
)

func SetupDB() *sql.DB {
	db := app.GetDB("root:mysql@tcp(localhost:3306)/belajar_golang_restful_api_test")
	return db
}

func SetupTest(db *sql.DB) *httprouter.Router {
	validate := validator.New()

	return routes.InitAllRoutes(db, validate)
}

func TruncateDB(db *sql.DB, tableName string) {
	_, err := db.ExecContext(context.Background(), "TRUNCATE "+tableName)
	helper.PanicIfError(err)
}
