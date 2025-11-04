package account

import (
	"encoding/json"
	"password/files"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	vault     *Vault
	vaultOnce sync.Once
)

func GetVault() *Vault {
	vaultOnce.Do(func() {
		vault = &Vault{
			UpdatedAt: time.Now(),
			Accounts:  []Account{},
		}
		db := files.NewJsonDB("data.json")
		data, err := db.Read()
		if err == nil {
			json.Unmarshal(data, vault)
		}
	})
	return vault
}

func (v *Vault) AddNewAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.save()
}

func (v *Vault) FindAccounts(url string) []*Account {
	res := []*Account{}
	for _, val := range v.Accounts {
		if strings.Contains(val.Url, url) {
			res = append(res, &val)
		}
	}
	return res
}

func (v *Vault) DeleteByUrl(url string) bool {
	res := false
	for i, val := range v.Accounts {
		if val.Url == url {
			v.Accounts = append(v.Accounts[:i], v.Accounts[i+1:]...)
			res = true
			v.save()
			break
		}
	}
	return res
}

func (v *Vault) ToByte() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Vault) save() {

	v.UpdatedAt = time.Now()
	data, err := v.ToByte()
	if err != nil {
		color.Red("Не удалось преобразовать в json")
	}
	db := files.NewJsonDB("data.json")
	db.Write(data)

}
