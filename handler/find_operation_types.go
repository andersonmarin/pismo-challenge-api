package handler

import (
	"net/http"

	"github.com/andersonmarin/pismo-challenge-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FindOperationTypes(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if operationTypes, err := usecase.FindOperationTypes(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			return ctx.JSON(http.StatusOK, &operationTypes)
		}
	}
}
