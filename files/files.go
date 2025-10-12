package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	NameDb string
}

func NewJsonDb() *JsonDb {
	name := "accounts.json"
	return &JsonDb{
		NameDb: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.NameDb)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {

	file, err := os.Create(db.NameDb)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись данных успешна.")
	defer file.Close()
}
