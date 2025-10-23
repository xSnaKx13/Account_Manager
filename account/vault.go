package account

import (
	"account-manager/files"
	promptdata "account-manager/promptData"
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accaunts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accaunts = append(vault.Accaunts, acc)
	vault.UpdatedAt = time.Now()
}

func (vault *VaultWithDb) SaveAndUpdate(db Db) {
	data, err := vault.ToBytes()
	if err != nil {
		promptdata.PrintError("Ошибка записи!")
	}
	db.Write(data)
}

func (vault *VaultWithDb) FindAccount(prompt string, checker func(Account, string) bool) {
	data := prompt
	var isFound bool
	for _, accounts := range vault.Accaunts {
		isFound = checker(accounts, data)
		if isFound {
			accounts.OutputAccount()
		}
	}
	if !isFound {
		fmt.Println("Аккаунт не найден!")
		return
	}
}

func (vault *VaultWithDb) DeleteAccount(checker func(Account, string) bool) {
	var foundAccounts []Account

	name := promptdata.PromptData("Введите логин (или часть логина) для удаления")
	var isFound bool
	for _, account := range vault.Accaunts {
		isFound = checker(account, name)
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
			promptdata.PrintError("Не удалось преобразовать в json!")
		}
		db := files.NewJsonDb()
		db.Write(data)
		return
	} else {
		promptdata.PrintError("Аккаунт не найден!")
		return
	}

}

func (vault *Vault) ToBytes() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		promptdata.PrintError("Ошибка преобразования в byte!")
		return nil, err
	}
	return data, nil
}

func NewVault(db Db) *VaultWithDb {

	data, err := db.Read()
	if err != nil {
		promptdata.PrintError("Файл отсутствует, будет создан новый файл!")
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
		promptdata.PrintError("Ошибка распаковки json!")
	}
	return &existingVault
}
