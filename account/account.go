package account

import (
	"account-manager/files"
	promptdata "account-manager/promptData"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `login:"login"`
	Password  string    `password:"password"`
	URL       string    `url:"url"`
	CreatedAt time.Time `created:"created"`
	UpdatedAt time.Time `url:"update"`
}

func (acc *Account) generatePassword(numberOfPasswordCharacters int) {
	acceptableСharacters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-_!№;%:?*()")
	pass := make([]rune, numberOfPasswordCharacters)
	for char := range pass {
		pass[char] = acceptableСharacters[rand.Intn(len(acceptableСharacters))]
	}
	acc.Password = string(pass)
}

func (acc *Account) OutputAccount() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Password)
	color.Magenta(acc.URL)
}

func CreateAccount() {
	err := newAccount()
	if err != nil {
		promptdata.PrintError("Ошибк создания аккаунта!")
		return
	}
	fmt.Print("Аккаунт создан!\n")
}

func newAccount() error {
	userLogin := promptdata.PromptData("Введите логин")
	if len(userLogin) == 0 {
		return errors.New("INVALID_LOGIN")
	}

	fmt.Println("~ пароль будет сгенерирован системой автоматически в слечае пропуска данного поля! ~")
	userPassword := promptdata.PromptData("Введите пароль: ")

	userUrl := promptdata.PromptData("Введите URL")
	_, err := url.ParseRequestURI(userUrl)
	if err != nil {
		return errors.New("INVALID_URL")
	}

	newAcc := &Account{
		Login:     userLogin,
		Password:  userPassword,
		URL:       userUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(newAcc.Password) == 0 {
		fmt.Println("Поле 'пароль' не может быть пустым.\nПароль сгенерирован автоматически!")
		newAcc.generatePassword(12)
	}
	db := files.NewJsonDb()
	vault := NewVault(db)
	vault.AddAccount(*newAcc)
	vault.SaveAndUpdate(db)
	return nil
}
func FindAccount() {
	vault := NewVault(files.NewJsonDb())
	vault.FindAccount()
}
func DeleteAccount() {
	vault := NewVault(files.NewJsonDb())
	vault.DeleteAccount()
}
