package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"restful-api/helper"
	"time"
)

func GetDB(dataSourceName ...string) *sql.DB {

	var customDataSourceName string

	if dataSourceName != nil {
		customDataSourceName = dataSourceName[0]
	} else {
		customDataSourceName = "root:mysql@tcp(localhost:3306)/belajar_golang_restful_api"
	}

	db, err := sql.Open("mysql", customDataSourceName)

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
