package services

import (
	"bank/models"
	"errors"
)

func Login(username, password string) (*models.Account, error) {
	accounts, _ := LoadAccounts()

	for _, acc := range accounts {
		if acc.Username == username && acc.Password == password {
			return &acc, nil
		}
	}

	return nil, errors.New("Invalid Username or Password!")
}
