package repository

import (
	"context"
	"database/sql"
	"errors"
	"restful-api/helper"
	"restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repo *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	SQL := "INSERT INTO category(name) VALUES (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category

}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category

}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {

	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {

	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	defer func() {
		err := rows.Close()
		helper.PanicIfError(err)
	}()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}

}

func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {

	SQL := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer func() {
		err := rows.Close()
		helper.PanicIfError(err)
	}()

	var categories []domain.Category
	for rows.Next() {

		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)

	}

	return categories

}
