package main

import (
	"account-manager/account"
	"bufio"
	"fmt"
	"os"
)

var menu = map[int]func(){
	1: account.CreateAccount,
	2: account.FindAccountByLogin,
	3: account.FindAccountByUrl,
	4: account.DeleteAccountByLogin,
}

func Menu() {
Menu:
	for {
		fmt.Println("Выберете действие: ")
		fmt.Println("1 - создать аккаунт")
		fmt.Println("2 - найти аккаунт по логину")
		fmt.Println("3 - найти аккаунт по URL")
		fmt.Println("4 - удалить аккаунт")
		fmt.Println("5 - выйти из программы")

		var input int
		fmt.Scan(&input)

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		menuFunc := menu[input]
		if menuFunc != nil {
			menuFunc()
		} else if input == 5 {
			break Menu
		} else {
			fmt.Println("Некорректный ввод, попробуйте еще раз!")
		}
	}

}
