package services

import (
	"bank/models"
	"encoding/json"
	"log"
	"os"
)

const dataFile = "data/users.json"

func LoadAccounts() ([]models.Account, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return []models.Account{}, nil
	}

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatalf("Can not read file: %v", err)
	}

	var accounts []models.Account
	err = json.Unmarshal(bytes, &accounts)

	return accounts, err
}

func SaveAccount(accounts []models.Account) error {
	bytes, err := json.MarshalIndent(accounts, "", "  ")

	if err != nil {
		log.Fatalf("Can not converted to JSON: %v", err)
	}

	return os.WriteFile(dataFile, bytes, 0777)
}
