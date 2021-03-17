package handler

import (
	"net/http"

	"github.com/andersonmarin/pismo-challenge-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTransaction(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var r struct {
			AccountID       uint64  `json:"account_id"`
			OperationTypeID uint64  `json:"operation_type_id"`
			Amount          float64 `json:"amount"`
		}
		if err := ctx.Bind(&r); err != nil {
			return err
		}

		if err := usecase.CheckTransactionAmount(db, r.OperationTypeID, r.Amount); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if transaction, err := usecase.CreateTransaction(db, r.AccountID, r.OperationTypeID, r.Amount); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			return ctx.JSON(http.StatusCreated, &transaction)
		}
	}
}
