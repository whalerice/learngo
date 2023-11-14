package main

import (
	"fmt"

	"github.com/learngo/accounts"
)

func main()  {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}