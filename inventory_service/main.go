package main

import (
	"net/http"
	"webstart/database"
	"webstart/product"
	"webstart/receipt"

	//underscore is used when we dont explicitly use library, but we need
	//it for side efects
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
