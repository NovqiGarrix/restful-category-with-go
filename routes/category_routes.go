package routes

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"restful-api/controller"
	"restful-api/repository"
	"restful-api/service"
)

func InitCategoryRoutes(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}
