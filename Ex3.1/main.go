package main

import (
	"bank/services"
	"fmt"
)

func main() {
	var user, pass string

	fmt.Print("Please input your Username: ")
	fmt.Scanln(&user)

	fmt.Print("Please input your Password: ")
	fmt.Scanln(&pass)

	account, err := services.Login(user, pass)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for {
		fmt.Println("\n1. Check Balancing")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Sent Money")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance: %.2f\n", account.Balance)
		case 2:
			var amount float64
			fmt.Printf("Please enter your deposit amount: ")
			fmt.Scanln(&amount)
			err := services.Deposit(account, amount)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			var amount float64
			fmt.Printf("Please enter your withdraw amount: ")
			fmt.Scanln(&amount)
			err := services.Withdraw(account, amount)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			var to string
			var amount float64

			fmt.Println("Receiver Account: ")
			fmt.Scanln(&to)
			fmt.Println("Transfer amount: ")
			fmt.Scanln(&amount)

			err := services.Transfer(account.Username, to, amount)

			if err != nil {
				fmt.Println("Error", err)
			} else {
				updateAccount, _ := services.Login(account.Username, account.Password)
				account = updateAccount
				fmt.Println("Money transfer successfully")
			}

		case 5:
			fmt.Println("Thank you for using out service!")
			return
		}
	}
}
