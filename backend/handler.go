package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/labstack/echo"
)

type HTTPMessageBody struct {
	Message string
	Time    string
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
		return ctx.JSON(http.StatusInternalServerError, HTTPMessageBody{
			Message: tmpErr.Error(), Time: time.Now().Format(time.DateTime)})
	}

	var vd VaultDto
	if bindErr := ctx.Bind(&vd); bindErr != nil {
		log.Println(bindErr)
		return ctx.JSON(http.StatusInternalServerError, HTTPMessageBody{
			Message: bindErr.Error(), Time: time.Now().Format(time.DateTime)})
	}

	if unlockErr := Unlock(vd.Unlock, vd.Url); unlockErr != nil {
		log.Println(unlockErr)
		return ctx.JSON(http.StatusInternalServerError, HTTPMessageBody{
			Message: unlockErr.Error(), Time: time.Now().Format(time.DateTime)})
	}

	items, getErr := GetItems(vd.Url)
	if getErr != nil {
		log.Println(getErr)
		return ctx.JSON(http.StatusInternalServerError, HTTPMessageBody{
			Message: getErr.Error(), Time: time.Now().Format(time.DateTime)})
	}

	file, createErr := os.Create("/vault/vault.env")
	if createErr != nil {
		log.Println(createErr)
		return ctx.JSON(http.StatusInternalServerError, HTTPMessageBody{
			Message: createErr.Error(), Time: time.Now().Format(time.DateTime)})
	}
	defer file.Close()

	if parseErr := t.Execute(file, GenerateFields(vd.FolderID, items)); parseErr != nil {
		log.Println(parseErr)
		return ctx.JSON(http.StatusInternalServerError, HTTPMessageBody{
			Message: parseErr.Error(), Time: time.Now().Format(time.DateTime)})
	}

	return ctx.JSON(http.StatusOK, HTTPMessageBody{
		Message: "vault set", Time: time.Now().Format(time.DateTime)})
}
