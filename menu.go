package main

import (
	"account-manager/account"
	"bufio"
	"fmt"
	"os"
)

func menu() {
Menu:
	for {
		fmt.Println("Выберете действие: ")
		fmt.Println("1 - создать аккаунт")
		fmt.Println("2 - найти аккаунт")
		fmt.Println("3 - удалить аккаунт")
		fmt.Println("4 - выход")

		var input int
		fmt.Scan(&input)

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		switch input {
		case 1:
			account.CreateAccount()
		case 2:
			account.FindAccount()
		case 3:
			account.DeleteAccount()
		case 4:
			break Menu
		default:
			fmt.Println("WRONG_INPUT")
		}
	}

}
