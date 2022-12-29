package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"restful-api/exception"
	"restful-api/helper"
	"restful-api/model/domain"
	"restful-api/model/domain/web"
	"restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, validate: validate}
}

func (c CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (c CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {

	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (c CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundException(err))
	}

	return helper.ToCategoryResponse(category)

}

func (c CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundException(err))
	}

	c.CategoryRepository.Delete(ctx, tx, category)

}

func (c CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)

}
