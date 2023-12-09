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
	Unlock   string `json:"unlock"`
	Url      string `json:"url"`
	FolderID string `json:"folder_id"`
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

	items, err := GetItems(vd.Url)
	if err != nil {
		return err
	}

	file, err := os.Create("/vault/vault.env")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	if parseErr := t.Execute(file, GenerateFields(vd.FolderID, items)); parseErr != nil {
		log.Println(parseErr)
		return parseErr
	}

	return ctx.JSON(http.StatusOK, HTTPMessageBody{Message: "vault set"})
}
