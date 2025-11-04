package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var Letters = [52]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (acc *Account) genPassword(length int) {
	arr := make([]int32, length)
	for i := 0; i < length; i += 1 {
		arr[i] = Letters[rand.IntN(len(Letters))]
	}
	acc.Password = string(arr)
}

func (acc *Account) Output() {
	color.Cyan("url: " + acc.Url)
	color.Cyan("login: " + acc.Login)
	color.Cyan("password: " + acc.Password)
}

func NewAccount(urlSting, password, login string) (*Account, error) {
	_, err := url.ParseRequestURI(urlSting)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	acc := &Account{
		Url:       urlSting,
		Login:     login,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		acc.genPassword(12)
	} else {
		acc.Password = password
	}
	return acc, nil
}
