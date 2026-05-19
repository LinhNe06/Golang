package services

import (
	"bank/models"
	"errors"
)

func UpdateAccount(updated models.Account) error {
	accounts, _ := LoadAccounts()
	for idx, acc := range accounts {
		if acc.Username == updated.Username {
			accounts[idx] = updated

			return SaveAccount(accounts)
		}
	}

	return errors.New("Account is not existed")
}

func Deposit(acc *models.Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be higher than zero")
	}

	acc.Balance += amount

	return UpdateAccount(*acc)
}

func Withdraw(acc *models.Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be higher than zero")
	}
	if amount > acc.Balance {
		return errors.New("Your balance is not enough")
	}

	acc.Balance -= amount
	return UpdateAccount(*acc)
}

func Transfer(from, toUser string, amount float64) error {
	accounts, _ := LoadAccounts()

	var sender *models.Account
	var receiver *models.Account

	for idx, acc := range accounts {
		if acc.Username == from {
			sender = &accounts[idx]
		}

		if acc.Username == toUser {
			receiver = &accounts[idx]
		}

	}
	if sender == nil || receiver == nil {
		return errors.New("Sender or Receiver account is not existed!")
	}

	if sender.Balance < amount || amount <= 0 {
		return errors.New("Not enough cash, please deposit more!")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	return SaveAccount(accounts)
}
