package account

import (
	"account-manager/files"
	promptdata "account-manager/promptData"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accaunts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated"`
}

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func (vault *VaultWithDb) AddAccount(acc Account, db *files.JsonDb) {
	vault.Accaunts = append(vault.Accaunts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	db.Write(data)
}

func (vault *VaultWithDb) FindAccount() {
	name := promptdata.PromptData("Введите логин (или часть логина) для поиска")
	var isFound bool
	for _, accounts := range vault.Accaunts {
		isFound = strings.Contains(accounts.Login, name)
		if isFound {
			fmt.Println(accounts.Login)
			fmt.Println(accounts.Password)
			fmt.Println(accounts.URL)
			fmt.Println(accounts.CreatedAt)
			return
		}
	}
	if !isFound {
		fmt.Println("Аккаунт не найден!")
		return
	}
}

func (vault *VaultWithDb) DeleteAccount() {
	var foundAccounts []Account

	name := promptdata.PromptData("Введите логин (или часть логина) для удаления")
	var isFound bool
	for _, account := range vault.Accaunts {
		isFound = strings.Contains(account.Login, name)
		if !isFound {
			foundAccounts = append(foundAccounts, account)
		}
	}
	vault.Accaunts = foundAccounts
	vault.UpdatedAt = time.Now()
	if isFound {
		fmt.Println("Аккаунт удален.")
		data, err := vault.ToBytes()
		if err != nil {
			fmt.Println("Не удалось преобразовать!")
		}
		db := files.NewJsonDb()
		db.Write(data)
		return
	} else {
		fmt.Println("Аккаунт не найден!")
		return
	}

}

func (vault *Vault) ToBytes() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func NewVault(db files.JsonDb) *VaultWithDb {

	data, err := db.Read()
	if err != nil {
		fmt.Println("Файл отсутствует, будет создан новый файл!")
		return &VaultWithDb{
			Vault: Vault{
				Accaunts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var existingVault VaultWithDb
	err = json.Unmarshal(data, &existingVault)
	if err != nil {
		color.Red(err.Error())
	}
	return &existingVault
}
