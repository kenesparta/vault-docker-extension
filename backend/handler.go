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

type VaultDto struct {
	Unlock string `json:"unlock"`
	Url    string `json:"url"`
}

func vault(ctx echo.Context) error {
	t, tmpErr := template.ParseFiles("./vars.tmpl")
	if tmpErr != nil {
		log.Println(tmpErr)
		return tmpErr
	}

	var vd VaultDto
	if bindErr := ctx.Bind(&vd); bindErr != nil {
		log.Println(bindErr)
		return bindErr
	}

	if unlockErr := Unlock(vd.Unlock, vd.Url); unlockErr != nil {
		log.Println(unlockErr)
		return unlockErr
	}

	data := Data{
		Name:  "Alice",
		Place: "Wonderland",
	}

	file, err := os.Create("/vault/vars.sh")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	if parseErr := t.Execute(file, data); parseErr != nil {
		log.Println(parseErr)
		return parseErr
	}

	return ctx.JSON(http.StatusOK, HTTPMessageBody{Message: "vault set"})
}
