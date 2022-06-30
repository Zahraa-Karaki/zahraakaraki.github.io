package handlers

import (
	"database/sql"
	"net/http"

	//"strconv"
	"github.com/labstack/echo"
	"github.com/zahraakaraki/MyApp/models"
)


func AddOperation(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var operation models.Operation

		c.Bind(&operation)

		answer, err := models.AddOperation(db, operation.FirstNumber, operation.SecondNumber, operation.Operator)

		operation.Answer = answer

		if err == nil {
			return c.JSON(http.StatusCreated, operation)
		} else {
			return c.JSON(http.StatusOK, operation)
		}
	}
}
