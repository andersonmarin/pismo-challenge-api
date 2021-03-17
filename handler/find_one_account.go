package handler

import (
	"net/http"
	"strconv"

	"github.com/andersonmarin/pismo-challenge-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FindOneAccount(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if account, err := usecase.FindOneAccount(db, id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			return ctx.JSON(http.StatusOK, &account)
		}
	}
}
