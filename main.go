package main

import (
	"fmt"
	"password/account"
	"password/files"
	"strconv"

	"github.com/fatih/color"
)

// func reverse(arr *[4]int) {
// 	lenngth := len(*arr)
// 	holf := lenngth / 2
// 	for index := 0; index < holf; index++ {
// 		(*arr)[lenngth-index-1], (*arr)[index] = (*arr)[index], (*arr)[lenngth-index-1]
// 	}
// }

func prompData(word string) string {
	var data string
	fmt.Println(word)
	fmt.Scanln(&data)
	return data
}

func addNewAccount() {
	var urlSting, password, login string
	fmt.Println("введите логин")
	fmt.Scanln(&login)
	fmt.Println("введите пароль")
	fmt.Scanln(&password)
	fmt.Println("введите url")
	fmt.Scanln(&urlSting)

	a, err := account.NewAccount(urlSting, password, login)
	if err != nil {
		fmt.Println(err)
	}
	v := account.GetVault()
	v.AddNewAccount(*a)
	color.Green("аккаунт добавлен")
}

func deleteAccount() {
	url := prompData("введите url")
	vault := account.GetVault()
	res := vault.DeleteByUrl(url)
	if res {
		color.Cyan("Аккаунт удален")
	} else {
		color.Red("Аккаунт не найден")
	}

}

func findAccount() {
	url := prompData("введите url")
	vault := account.GetVault()
	accs := vault.FindAccounts(url)
	if len(accs) == 0 {
		color.Red("аккаунтов не найдено")
	} else {

		fmt.Println("найдено : " + strconv.Itoa(len(accs)))
		for i, acc := range accs {
			color.Cyan("аккаунт: " + strconv.Itoa(i+1))
			acc.Output()
			color.Cyan("_________")
		}
	}
}

func showMenu() {
	fmt.Println("1: создать аккаунт")
	fmt.Println("2: удалить аккаунт")
	fmt.Println("3: найти аккаунт")
	fmt.Println("4: вызод")
}

func main() {

	var choise int
Menu:
	for {
		showMenu()
		fmt.Println("выберете действие: ")
		fmt.Scanln(&choise)
		switch choise {

		case 1:
			addNewAccount()
		case 2:
			deleteAccount()
		case 3:
			findAccount()
		default:
			break Menu
		}

	}
	vault := account.GetVault()
	files.ReadFile("somePath")
	data, _ := vault.ToByte()
	files.WriteFile("data.json", data)
}

// func main() {

// 	forc := map[int]func(){
// 		1: addNewAccount,
// 		2: deleteAccount,
// 		3: findAccount,
// 	}

// 	var choise int
// 	var choiseFun func()
// 	for {
// 		showMenu()
// 		fmt.Println("выберете действие: ")
// 		fmt.Scanln(&choise)
// 		if choise == 0 {
// 			break
// 		}

// 		choiseFun = forc[choise]
// 		if choiseFun != nil {
// 			choiseFun()
// 		}
// 	}

// 	files.ReadFile("somePath")
// 	files.WriteFile("test.txt", "someData")
// }
