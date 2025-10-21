package main

import (
	"account-manager/account"
	"bufio"
	"fmt"
	"os"
)

var menu = map[int]func(){
	1: account.CreateAccount,
	2: account.FindAccount,
	3: account.DeleteAccount,
}

func Menu() {
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

		menuFunc := menu[input]
		if menuFunc != nil {
			menuFunc()
		} else if input == 4 {
			break Menu
		} else {
			fmt.Println("Некорректный ввод, попробуйте еще раз!")
		}
	}

}
