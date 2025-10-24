package main

import (
	promptdata "account-manager/promptData"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		promptdata.PrintError("Ошибка загрузки .env файла")
	}
	key := os.Getenv("KEY")
	password := os.Getenv("PASSWORD")
	fmt.Printf("KEY: %s\n", key)
	fmt.Printf("PASSWORD: %s\n", password)
	Menu()
}
