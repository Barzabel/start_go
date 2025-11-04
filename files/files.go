package files

import (
	"os"

	"github.com/fatih/color"
)

type JsonDB struct {
	name string
}

func NewJsonDB(name string) *JsonDB {
	res := JsonDB{
		name: name,
	}
	return &res
}

func (db *JsonDB) Read() ([]byte, error) {
	file, err := os.ReadFile(db.name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (db *JsonDB) Write(data []byte) {
	file, err := os.Create(db.name)
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		color.Red(err.Error())
	}
}
