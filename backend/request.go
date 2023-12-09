package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	req, err := http.NewRequest("POST", url+"/unlock", bytes.NewBuffer(jsonData))
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
	if resp.StatusCode != 200 {
		return errors.New("request error: " + string(body))
	}

	return nil
}

func GetItems(url string) ([]VaultItem, error) {
	req, err := http.NewRequest("GET", url+"/list/object/items", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var dataResult DataPayload
	json.Unmarshal(body, &dataResult)

	fmt.Println("Response:", string(body))
	return dataResult.Data.Data, nil
}
