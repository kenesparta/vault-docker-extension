package main

import (
	"time"
)

type DataPayload struct {
	Success bool `json:"success"`
	Data    struct {
		Object string      `json:"object"`
		Data   []VaultItem `json:"data"`
	} `json:"data"`
}

type VaultItem struct {
	PasswordHistory interface{} `json:"passwordHistory"`
	RevisionDate    time.Time   `json:"revisionDate"`
	CreationDate    time.Time   `json:"creationDate"`
	DeletedDate     interface{} `json:"deletedDate"`
	Object          string      `json:"object"`
	Id              string      `json:"id"`
	OrganizationId  interface{} `json:"organizationId"`
	FolderId        string      `json:"folderId"`
	Type            int         `json:"type"`
	Reprompt        int         `json:"reprompt"`
	Name            string      `json:"name"`
	Notes           interface{} `json:"notes"`
	Favorite        bool        `json:"favorite"`
	Fields          []struct {
		Name     string      `json:"name"`
		Value    string      `json:"value"`
		Type     int         `json:"type"`
		LinkedId interface{} `json:"linkedId"`
	} `json:"fields"`
	Login struct {
		Fido2Credentials     []interface{} `json:"fido2Credentials"`
		Username             string        `json:"username"`
		Password             string        `json:"password"`
		Totp                 interface{}   `json:"totp"`
		PasswordRevisionDate interface{}   `json:"passwordRevisionDate"`
	} `json:"login"`
	CollectionIds []interface{} `json:"collectionIds"`
}

type Field struct {
	Name  string
	Value string
}

func GenerateFields(folderId string, vaultItems []VaultItem) []Field {
	var fields []Field
	for _, vi := range vaultItems {
		if folderId != vi.FolderId {
			continue
		}
		for _, f := range vi.Fields {
			fields = append(fields, Field{
				Name:  f.Name,
				Value: f.Value,
			})
		}
	}
	return fields
}
