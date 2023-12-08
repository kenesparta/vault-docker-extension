package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Unlock(password, url string) error {
	payload := map[string]string{"password": password}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Response:", string(body))

	return nil
}
