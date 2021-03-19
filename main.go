package main

import (
	"os"

	"github.com/andersonmarin/pismo-challenge-api/handler"
	"github.com/andersonmarin/pismo-challenge-api/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := infrastructure.Database()

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Gzip(), middleware.CORS())

	e.POST("/accounts", handler.CreateAccount(db))
	e.GET("/accounts/:id", handler.FindOneAccount(db))
	e.GET("/operation-types", handler.FindOperationTypes(db))
	e.POST("/transactions", handler.CreateTransaction(db))

	if err := e.Start(os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
