package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zahraakaraki/MyApp/models"
)

type J map[string]interface{}

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

func GetOperations(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, models.GetOperations(db))
	}
}

func DeleteOperation(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := models.DeleteOperation(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, J{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
