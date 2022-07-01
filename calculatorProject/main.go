package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/zahraakaraki/MyApp/handlers"
)

func main() {

	var err error
	db, err := sql.Open("postgres", "user=postgres password=karaki217 dbname=calculator sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	e.File("/", "frontend/index.html")

	e.GET("/operations", handlers.GetOperations(db))

	e.POST("/operations", handlers.AddOperation(db))

	e.DELETE("/operations/:id", handlers.DeleteOperation(db))

	e.Start(":8080")

}
