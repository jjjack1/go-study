package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Account struct {
	balance int
}

func (a *Account) Withdraw(val int) {
	a.balance -= val
}
func (a *Account) Deposit(val int) {
	a.balance += val
}

func (a *Account) Balance() int {
	return a.balance
}

var accounts []*Account

func Transfer(sender, receiver int, money int) {
	accounts[sender].Withdraw(money)
	accounts[receiver].Deposit(money)
}

func GetTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()

	}
	return total
}

func Randomtransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)

}

func GoTransfer() {
	for {
		Randomtransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d \n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000})

	}
	PrintTotalBalance()
	for i := 0; i < 10; i++ {
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
