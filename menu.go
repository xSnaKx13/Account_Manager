package main

import (
	"account-manager/account"
	promptdata "account-manager/promptData"
	"fmt"
)

var menu = map[string]func(){
	"1": account.CreateAccount,
	"2": account.FindAccountByLogin,
	"3": account.FindAccountByUrl,
	"4": account.DeleteAccountByLogin,
}

func Menu() {
Menu:
	for {
		input := promptdata.PromptData(
			"Выберете действие: \n",
			"1 - создать аккаунт\n",
			"2 - найти аккаунт по логину\n",
			"3 - найти аккаунт по URL\n",
			"4 - удалить аккаунт\n",
			"5 - выйти из программы\n",
		)

		menuFunc := menu[input]
		if menuFunc != nil {
			menuFunc()
		} else if input == "5" {
			break Menu
		} else {
			fmt.Println("Некорректный ввод, попробуйте еще раз!")
		}
	}

}
