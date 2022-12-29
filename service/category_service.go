package service

import (
	"context"
	"restful-api/model/domain/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindAll(ctx context.Context) []web.CategoryResponse
}
