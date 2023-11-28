package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/labstack/echo"
)

type HTTPMessageBody struct {
	Message string
}

type Data struct {
	Name  string
	Place string
}

func vault(ctx echo.Context) error {
	t, err := template.ParseFiles("./vars.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	data := Data{
		Name:  "Alice",
		Place: "Wonderland",
	}

	file, err := os.Create("/vault/vars.sh")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		log.Fatal(err)
	}

	return ctx.JSON(http.StatusOK, HTTPMessageBody{Message: "vault set"})
}
