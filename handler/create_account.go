package handler

import (
	"net/http"

	"github.com/andersonmarin/pismo-challenge-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var r struct {
			DocumentNumber string `json:"document_number"`
		}
		if err := ctx.Bind(&r); err != nil {
			return err
		}

		if err := usecase.CheckAccountDocumentNumber(r.DocumentNumber); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := usecase.CheckAccountAlreadyExists(db, r.DocumentNumber); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if account, err := usecase.CreateAccount(db, r.DocumentNumber); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			return ctx.JSON(http.StatusCreated, &account)
		}
	}
}
